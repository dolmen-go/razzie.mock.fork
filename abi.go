package mock

import (
	"reflect"
	"unsafe"

	"golang.org/x/exp/constraints"
)

//go:linkname addReflectOff reflect.addReflectOff
func addReflectOff(ptr unsafe.Pointer) int32

const (
	KindDirectIface uint8 = 1 << 5
)

type abiType struct {
	Size_       uintptr
	PtrBytes    uintptr // number of (prefix) bytes in the type that can contain pointers
	Hash        uint32  // hash of type; avoids computation in hash tables
	TFlag       uint8   // extra type information flags
	Align_      uint8   // alignment of variable with this type
	FieldAlign_ uint8   // alignment of struct field with this type
	Kind_       uint8   // enumeration for C
	Equal       func(unsafe.Pointer, unsafe.Pointer) bool
	GCData      *byte
	Str         int32 // string form
	PtrToThis   int32 // type for pointer to this type, may be zero
}

func toAbiType(typ reflect.Type) *abiType {
	type iface struct {
		typ  *abiType
		word unsafe.Pointer
	}
	t := (*iface)(unsafe.Pointer(&typ)).word
	return (*abiType)(t)
}

func (t *abiType) Pointers() bool {
	return t.PtrBytes != 0
}

func (t *abiType) IfaceIndir() bool {
	return t.Kind_&KindDirectIface == 0
}

type abiMethod struct {
	Name int32 // name of method
	Mtyp int32 // method type (without receiver)
	Ifn  int32 // fn used in interface call (one-word receiver)
	Tfn  int32 // fn used for normal method call
}

func (m *abiMethod) SetFn(fn any) {
	off := addReflectOff(unsafe.Pointer(reflect.ValueOf(fn).Pointer()))
	//m.Tfn = off
	m.Ifn = off
}

func getAbiMethods(typ reflect.Type) []abiMethod {
	if typ.Kind() != reflect.Struct {
		panic("not struct kind: " + typ.Kind().String())
	}
	type UncommonType struct {
		PkgPath int32  // import path; empty for built-in types like int, string
		Mcount  uint16 // number of methods
		Xcount  uint16 // number of exported methods
		Moff    uint32 // offset from this uncommontype to [mcount]Method
		_       uint32 // unused
	}
	type Name struct {
		Bytes *byte
	}
	type StructField struct {
		Name   Name     // name is always non-empty
		Typ    *abiType // type of field
		Offset uintptr  // byte offset of field
	}
	type StructType struct {
		abiType
		PkgPath Name
		Fields  []StructField
	}
	type StructTypeUncommon struct {
		StructType
		u UncommonType
	}
	t := toAbiType(typ)
	u := &(*StructTypeUncommon)(unsafe.Pointer(t)).u
	if u.Mcount == 0 {
		return nil
	}
	m := unsafe.Add(unsafe.Pointer(u), u.Moff)
	return ptrToSlice[abiMethod](m, u.Mcount)
}

func ptrToSlice[T any, N constraints.Integer](ptr unsafe.Pointer, len N) []T {
	return unsafe.Slice((*T)(ptr), len)
}

// abiStep represents an ABI "instruction." Each instruction
// describes one part of how to translate between a Go value
// in memory and a call frame.
type abiStep struct {
	kind abiStepKind

	// offset and size together describe a part of a Go value
	// in memory.
	offset uintptr
	size   uintptr // size in bytes of the part

	// These fields describe the ABI side of the translation.
	stkOff uintptr // stack offset, used if kind == abiStepStack
	ireg   int     // integer register index, used if kind == abiStepIntReg or kind == abiStepPointer
	freg   int     // FP register index, used if kind == abiStepFloatReg
}

// abiStepKind is the "op-code" for an abiStep instruction.
type abiStepKind int

const (
	abiStepBad      abiStepKind = iota
	abiStepStack                // copy to/from stack
	abiStepIntReg               // copy to/from integer register
	abiStepPointer              // copy pointer to/from integer register
	abiStepFloatReg             // copy to/from FP register
)

// abiSeq represents a sequence of ABI instructions for copying
// from a series of reflect.Values to a call frame (for call arguments)
// or vice-versa (for call results).
//
// An abiSeq should be populated by calling its addArg method.
type abiSeq struct {
	// steps is the set of instructions.
	//
	// The instructions are grouped together by whole arguments,
	// with the starting index for the instructions
	// of the i'th Go value available in valueStart.
	//
	// For instance, if this abiSeq represents 3 arguments
	// passed to a function, then the 2nd argument's steps
	// begin at steps[valueStart[1]].
	//
	// Because reflect accepts Go arguments in distinct
	// Values and each Value is stored separately, each abiStep
	// that begins a new argument will have its offset
	// field == 0.
	steps      []abiStep
	valueStart []int

	stackBytes   uintptr // stack space used
	iregs, fregs int     // registers used
}

// stepsForValue returns the ABI instructions for translating
// the i'th Go argument or return value represented by this
// abiSeq to the Go ABI.
func (a *abiSeq) stepsForValue(i int) []abiStep {
	s := a.valueStart[i]
	var e int
	if i == len(a.valueStart)-1 {
		e = len(a.steps)
	} else {
		e = a.valueStart[i+1]
	}
	return a.steps[s:e]
}

// addArg extends the abiSeq with a new Go value of type t.
//
// If the value was stack-assigned, returns the single
// abiStep describing that translation, and nil otherwise.
func (a *abiSeq) addArg(t reflect.Type) *abiStep {
	// We'll always be adding a new value, so do that first.
	pStart := len(a.steps)
	a.valueStart = append(a.valueStart, pStart)
	if t.Size() == 0 {
		// If the size of the argument type is zero, then
		// in order to degrade gracefully into ABI0, we need
		// to stack-assign this type. The reason is that
		// although zero-sized types take up no space on the
		// stack, they do cause the next argument to be aligned.
		// So just do that here, but don't bother actually
		// generating a new ABI step for it (there's nothing to
		// actually copy).
		//
		// We cannot handle this in the recursive case of
		// regAssign because zero-sized *fields* of a
		// non-zero-sized struct do not cause it to be
		// stack-assigned. So we need a special case here
		// at the top.
		a.stackBytes = align(a.stackBytes, uintptr(t.Align()))
		return nil
	}
	// Hold a copy of "a" so that we can roll back if
	// register assignment fails.
	aOld := *a
	if !a.regAssign(t, 0) {
		// Register assignment failed. Roll back any changes
		// and stack-assign.
		*a = aOld
		a.stackAssign(t.Size(), uintptr(t.Align()))
		return &a.steps[len(a.steps)-1]
	}
	return nil
}

// addRcvr extends the abiSeq with a new method call
// receiver according to the interface calling convention.
//
// If the receiver was stack-assigned, returns the single
// abiStep describing that translation, and nil otherwise.
// Returns true if the receiver is a pointer.
func (a *abiSeq) addRcvr(rcvr reflect.Type) (*abiStep, bool) {
	// The receiver is always one word.
	a.valueStart = append(a.valueStart, len(a.steps))
	var ok, ptr bool
	rcvr_ := toAbiType(rcvr)
	if rcvr_.IfaceIndir() || rcvr_.Pointers() {
		ok = a.assignIntN(0, ptrSize, 1, 0b1)
		ptr = true
	} else {
		// TODO(mknyszek): Is this case even possible?
		// The interface data work never contains a non-pointer
		// value. This case was copied over from older code
		// in the reflect package which only conditionally added
		// a pointer bit to the reflect.(Value).Call stack frame's
		// GC bitmap.
		ok = a.assignIntN(0, ptrSize, 1, 0b0)
		ptr = false
	}
	if !ok {
		a.stackAssign(ptrSize, ptrSize)
		return &a.steps[len(a.steps)-1], ptr
	}
	return nil, ptr
}

// regAssign attempts to reserve argument registers for a value of
// type t, stored at some offset.
//
// It returns whether or not the assignment succeeded, but
// leaves any changes it made to a.steps behind, so the caller
// must undo that work by adjusting a.steps if it fails.
//
// This method along with the assign* methods represent the
// complete register-assignment algorithm for the Go ABI.
func (a *abiSeq) regAssign(t reflect.Type, offset uintptr) bool {
	switch t.Kind() {
	case reflect.UnsafePointer, reflect.Pointer, reflect.Chan, reflect.Map, reflect.Func:
		return a.assignIntN(offset, t.Size(), 1, 0b1)
	case reflect.Bool, reflect.Int, reflect.Uint, reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Uintptr:
		return a.assignIntN(offset, t.Size(), 1, 0b0)
	case reflect.Int64, reflect.Uint64:
		switch ptrSize {
		case 4:
			return a.assignIntN(offset, 4, 2, 0b0)
		case 8:
			return a.assignIntN(offset, 8, 1, 0b0)
		}
	case reflect.Float32, reflect.Float64:
		return a.assignFloatN(offset, t.Size(), 1)
	case reflect.Complex64:
		return a.assignFloatN(offset, 4, 2)
	case reflect.Complex128:
		return a.assignFloatN(offset, 8, 2)
	case reflect.String:
		return a.assignIntN(offset, ptrSize, 2, 0b01)
	case reflect.Interface:
		return a.assignIntN(offset, ptrSize, 2, 0b10)
	case reflect.Slice:
		return a.assignIntN(offset, ptrSize, 3, 0b001)
	case reflect.Array:
		switch t.Len() {
		case 0:
			// There's nothing to assign, so don't modify
			// a.steps but succeed so the caller doesn't
			// try to stack-assign this value.
			return true
		case 1:
			return a.regAssign(t.Elem(), offset)
		default:
			return false
		}
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if !a.regAssign(f.Type, offset+f.Offset) {
				return false
			}
		}
		return true
	default:
		print("t.Kind == ", t.Kind(), "\n")
		panic("unknown type kind")
	}
	panic("unhandled register assignment path")
}

// assignIntN assigns n values to registers, each "size" bytes large,
// from the data at [offset, offset+n*size) in memory. Each value at
// [offset+i*size, offset+(i+1)*size) for i < n is assigned to the
// next n integer registers.
//
// Bit i in ptrMap indicates whether the i'th value is a pointer.
// n must be <= 8.
//
// Returns whether assignment succeeded.
func (a *abiSeq) assignIntN(offset, size uintptr, n int, ptrMap uint8) bool {
	if n > 8 || n < 0 {
		panic("invalid n")
	}
	if ptrMap != 0 && size != ptrSize {
		panic("non-empty pointer map passed for non-pointer-size values")
	}
	if a.iregs+n > intArgRegs {
		return false
	}
	for i := 0; i < n; i++ {
		kind := abiStepIntReg
		if ptrMap&(uint8(1)<<i) != 0 {
			kind = abiStepPointer
		}
		a.steps = append(a.steps, abiStep{
			kind:   kind,
			offset: offset + uintptr(i)*size,
			size:   size,
			ireg:   a.iregs,
		})
		a.iregs++
	}
	return true
}

// assignFloatN assigns n values to registers, each "size" bytes large,
// from the data at [offset, offset+n*size) in memory. Each value at
// [offset+i*size, offset+(i+1)*size) for i < n is assigned to the
// next n floating-point registers.
//
// Returns whether assignment succeeded.
func (a *abiSeq) assignFloatN(offset, size uintptr, n int) bool {
	if n < 0 {
		panic("invalid n")
	}
	if a.fregs+n > floatArgRegs || floatRegSize < size {
		return false
	}
	for i := 0; i < n; i++ {
		a.steps = append(a.steps, abiStep{
			kind:   abiStepFloatReg,
			offset: offset + uintptr(i)*size,
			size:   size,
			freg:   a.fregs,
		})
		a.fregs++
	}
	return true
}

// stackAssign reserves space for one value that is "size" bytes
// large with alignment "alignment" to the stack.
//
// Should not be called directly; use addArg instead.
func (a *abiSeq) stackAssign(size, alignment uintptr) {
	a.stackBytes = align(a.stackBytes, alignment)
	a.steps = append(a.steps, abiStep{
		kind:   abiStepStack,
		offset: 0, // Only used for whole arguments, so the memory offset is 0.
		size:   size,
		stkOff: a.stackBytes,
	})
	a.stackBytes += size
}

// abiDesc describes the ABI for a function or method.
type abiDesc struct {
	// call and ret represent the translation steps for
	// the call and return paths of a Go function.
	call, ret abiSeq

	// These fields describe the stack space allocated
	// for the call. stackCallArgsSize is the amount of space
	// reserved for arguments but not return values. retOffset
	// is the offset at which return values begin, and
	// spill is the size in bytes of additional space reserved
	// to spill argument registers into in case of preemption in
	// reflectcall's stack frame.
	stackCallArgsSize, retOffset, spill uintptr

	// stackPtrs is a bitmap that indicates whether
	// each word in the ABI stack space (stack-assigned
	// args + return values) is a pointer. Used
	// as the heap pointer bitmap for stack space
	// passed to reflectcall.
	stackPtrs *bitVector

	// inRegPtrs is a bitmap whose i'th bit indicates
	// whether the i'th integer argument register contains
	// a pointer. Used by makeFuncStub and methodValueCall
	// to make result pointers visible to the GC.
	//
	// outRegPtrs is the same, but for result values.
	// Used by reflectcall to make result pointers visible
	// to the GC.
	inRegPtrs, outRegPtrs intArgRegBitmap
}

func newAbiDesc(t reflect.Type, rcvr reflect.Type) abiDesc {
	if t.Kind() != reflect.Func {
		panic("not a function: " + t.Kind().String())
	}
	// We need to add space for this argument to
	// the frame so that it can spill args into it.
	//
	// The size of this space is just the sum of the sizes
	// of each register-allocated type.
	//
	// TODO(mknyszek): Remove this when we no longer have
	// caller reserved spill space.
	spill := uintptr(0)

	// Compute gc program & stack bitmap for stack arguments
	stackPtrs := new(bitVector)

	// Compute the stack frame pointer bitmap and register
	// pointer bitmap for arguments.
	inRegPtrs := intArgRegBitmap{}

	// Compute abiSeq for input parameters.
	var in abiSeq
	if rcvr != nil {
		stkStep, isPtr := in.addRcvr(rcvr)
		if stkStep != nil {
			if isPtr {
				stackPtrs.append(1)
			} else {
				stackPtrs.append(0)
			}
		} else {
			spill += ptrSize
		}
	}
	for i := 0; i < t.NumIn(); i++ {
		arg := t.In(i)
		stkStep := in.addArg(arg)
		if stkStep != nil {
			addTypeBits(stackPtrs, stkStep.stkOff, arg)
		} else {
			spill = align(spill, uintptr(arg.Align()))
			spill += arg.Size()
			for _, st := range in.stepsForValue(i) {
				if st.kind == abiStepPointer {
					inRegPtrs.Set(st.ireg)
				}
			}
		}
	}
	spill = align(spill, ptrSize)

	// From the input parameters alone, we now know
	// the stackCallArgsSize and retOffset.
	stackCallArgsSize := in.stackBytes
	retOffset := align(in.stackBytes, ptrSize)

	// Compute the stack frame pointer bitmap and register
	// pointer bitmap for return values.
	outRegPtrs := intArgRegBitmap{}

	// Compute abiSeq for output parameters.
	var out abiSeq
	// Stack-assigned return values do not share
	// space with arguments like they do with registers,
	// so we need to inject a stack offset here.
	// Fake it by artificially extending stackBytes by
	// the return offset.
	out.stackBytes = retOffset
	for i := 0; i < t.NumOut(); i++ {
		res := t.Out(i)
		stkStep := out.addArg(res)
		if stkStep != nil {
			addTypeBits(stackPtrs, stkStep.stkOff, res)
		} else {
			for _, st := range out.stepsForValue(i) {
				if st.kind == abiStepPointer {
					outRegPtrs.Set(st.ireg)
				}
			}
		}
	}
	// Undo the faking from earlier so that stackBytes
	// is accurate.
	out.stackBytes -= retOffset
	return abiDesc{in, out, stackCallArgsSize, retOffset, spill, stackPtrs, inRegPtrs, outRegPtrs}
}

// align returns the result of rounding x up to a multiple of n.
// n must be a power of two.
func align(x, n uintptr) uintptr {
	return (x + n - 1) &^ (n - 1)
}

// Note: this type must agree with runtime.bitvector.
type bitVector struct {
	n    uint32 // number of bits
	data []byte
}

// append a bit to the bitmap.
func (bv *bitVector) append(bit uint8) {
	if bv.n%(8*ptrSize) == 0 {
		// Runtime needs pointer masks to be a multiple of uintptr in size.
		// Since reflect passes bv.data directly to the runtime as a pointer mask,
		// we append a full uintptr of zeros at a time.
		for i := 0; i < ptrSize; i++ {
			bv.data = append(bv.data, 0)
		}
	}
	bv.data[bv.n/8] |= bit << (bv.n % 8)
	bv.n++
}

func addTypeBits(bv *bitVector, offset uintptr, t reflect.Type) {
	if !toAbiType(t).Pointers() {
		return
	}

	switch t.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.Slice, reflect.String, reflect.UnsafePointer:
		// 1 pointer at start of representation
		for bv.n < uint32(offset/ptrSize) {
			bv.append(0)
		}
		bv.append(1)

	case reflect.Interface:
		// 2 pointers
		for bv.n < uint32(offset/ptrSize) {
			bv.append(0)
		}
		bv.append(1)
		bv.append(1)

	case reflect.Array:
		// repeat inner type
		for i := 0; i < t.Len(); i++ {
			addTypeBits(bv, offset+uintptr(i)*t.Elem().Size(), t.Elem())
		}

	case reflect.Struct:
		// apply fields
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			addTypeBits(bv, offset+f.Offset, f.Type)
		}
	}
}

// intArgRegBitmap is a bitmap large enough to hold one bit per
// integer argument/return register.
type intArgRegBitmap [(intArgRegs + 7) / 8]uint8

// Set sets the i'th bit of the bitmap to 1.
func (b *intArgRegBitmap) Set(i int) {
	b[i/8] |= uint8(1) << (i % 8)
}

// Get returns whether the i'th bit of the bitmap is set.
//
// nosplit because it's called in extremely sensitive contexts, like
// on the reflectcall return path.
//
//go:nosplit
func (b *intArgRegBitmap) Get(i int) bool {
	return b[i/8]&(uint8(1)<<(i%8)) != 0
}
