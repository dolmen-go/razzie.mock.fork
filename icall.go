package mock

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/stretchr/testify/mock"
)

type icallBase []*methodInfo

type methodInfo struct {
	typ      reflect.Type
	name     string
	abiDesc  abiDesc
	inTypes  []reflect.Type
	outTypes []reflect.Type
}

func newMethodInfo(typ reflect.Type, method reflect.Method) *methodInfo {
	inTypes, outTypes := getMethodInOutTypes(method)
	return &methodInfo{
		typ:      typ,
		name:     method.Name,
		abiDesc:  newAbiDesc(method.Type, typ),
		inTypes:  inTypes,
		outTypes: outTypes,
	}
}

func icall(index int, recv unsafe.Pointer, intRegs []unsafe.Pointer, floatRegs []float64, frame unsafe.Pointer) (returnIntRegs [9]unsafe.Pointer, returnFloatRegs [15]float64) {
	receiver := reflect.NewAt(reflect.TypeFor[icallBase](), recv).Elem()
	info := receiver.Interface().(icallBase)[index]
	receiver = reflect.NewAt(info.typ, recv).Elem()
	mock := receiver.Field(1).Interface().(*mock.Mock)
	ptr := frame
	var in []any
	for i, typ := range info.inTypes {
		if typ.Size() == 0 {
			in = append(in, reflect.Zero(typ).Interface())
			continue
		}
		var v reflect.Value
		steps := info.abiDesc.call.stepsForValue(i + 1) // skip receiver
		if st := steps[0]; st.kind == abiStepStack {
			if toAbiType(typ).IfaceIndir() {
				// value cannot be inlined in interface data.
				// Must make a copy, because f might keep a reference to it,
				// and we cannot let f keep a reference to the stack frame
				// after this function returns, not even a read-only reference.
				v = reflect.New(typ).Elem()
				v_ := reflect.NewAt(typ, unsafe.Add(ptr, st.stkOff)).Elem()
				v.Set(v_)
			} else {
				v = reflect.NewAt(typ, *(*unsafe.Pointer)(unsafe.Add(ptr, st.stkOff))).Elem()
			}
		} else {
			if toAbiType(typ).IfaceIndir() {
				// All that's left is values passed in registers that we need to
				// create space for the values.
				vptr := unsafe_New(typ)
				v = reflect.NewAt(typ, vptr).Elem()
				for _, st := range steps {
					switch st.kind {
					case abiStepIntReg:
						offset := unsafe.Add(vptr, st.offset)
						regPtr := unsafe.Pointer(&intRegs[st.ireg])
						if bigEndian {
							regPtr = unsafe.Add(regPtr, ptrSize-st.size)
						}
						memmove(offset, regPtr, st.size)
					case abiStepPointer:
						s := unsafe.Add(vptr, st.offset)
						*((*unsafe.Pointer)(s)) = intRegs[st.ireg]
					case abiStepFloatReg:
						offset := unsafe.Add(vptr, st.offset)
						switch st.size {
						case 4:
							*(*float32)(offset) = *(*float32)(unsafe.Pointer(&floatRegs[st.freg]))
						case 8:
							*(*float64)(offset) = floatRegs[st.freg]
						default:
							panic("bad st.size")
						}
					case abiStepStack:
						panic("register-based return value has stack component")
					default:
						panic("unknown ABI part kind")
					}
				}
			} else {
				// Pointer-valued data gets put directly
				// into v.ptr.
				if steps[0].kind != abiStepPointer {
					print("kind=", steps[0].kind, ", type=", typ.Name(), "\n")
					panic("mismatch between ABI description and types")
				}
				v = reflect.NewAt(typ, unsafe.Pointer(&intRegs[steps[0].ireg])).Elem()
			}
		}
		if v.IsZero() {
			in = append(in, nil)
		} else {
			in = append(in, v.Interface())
		}
	}
	out := mock.MethodCalled(info.name, in...)
	if len(out) != len(info.outTypes) {
		panic("mock: wrong return count from function: " + info.name)
	}
	for i, typ := range info.outTypes {
		if typ.Size() == 0 {
			continue
		}
		vptr := unsafe_New(typ)
		v := reflect.NewAt(typ, vptr).Elem()
		if out[i] != nil {
			v.Set(reflect.ValueOf(out[i]))
		}
		steps := info.abiDesc.ret.stepsForValue(i)
		if st := steps[0]; st.kind == abiStepStack {
			// Copy values to the "stack."
			addr := unsafe.Add(ptr, st.stkOff)
			// Do not use write barriers. The stack space used
			// for this call is not adequately zeroed, and we
			// are careful to keep the arguments alive until we
			// return to makeFuncStub's caller.
			if toAbiType(typ).IfaceIndir() {
				memmove(addr, vptr, st.size)
			} else {
				*(*unsafe.Pointer)(addr) = vptr
			}
		} else {
			for _, st := range steps {
				switch st.kind {
				case abiStepIntReg, abiStepPointer:
					// Copy values to "integer registers."
					if toAbiType(typ).IfaceIndir() {
						offset := unsafe.Add(vptr, st.offset)
						regPtr := unsafe.Pointer(&returnIntRegs[st.ireg])
						if bigEndian {
							regPtr = unsafe.Add(regPtr, ptrSize-st.size)
						}
						memmove(regPtr, offset, st.size)
					} else {
						// Only populate the Ints space on the return path.
						// This is safe because out is kept alive until the
						// end of this function, and the return path through
						// makeFuncStub has no preemption, so these pointers
						// are always visible to the GC.
						returnIntRegs[st.ireg] = *(*unsafe.Pointer)(vptr)
					}
				case abiStepFloatReg:
					// Copy values to "float registers."
					if !toAbiType(typ).IfaceIndir() {
						panic("attempted to copy pointer to FP register")
					}
					offset := unsafe.Add(vptr, st.offset)
					switch st.size {
					case 4:
						*(*float32)(unsafe.Pointer(&returnFloatRegs[st.freg])) = *(*float32)(offset)
					case 8:
						returnFloatRegs[st.freg] = *(*float64)(offset)
					default:
						panic("bad st.size")
					}
				default:
					panic("unknown ABI part kind")

				}
			}
		}
	}
	runtime.KeepAlive(out)
	return
}

func getMethodInOutTypes(method reflect.Method) (in, out []reflect.Type) {
	typ := method.Type
	for i := 0; i < typ.NumIn(); i++ {
		in = append(in, typ.In(i))
	}
	for i := 0; i < typ.NumOut(); i++ {
		out = append(out, typ.Out(i))
	}
	return
}

func memmove(dst, src unsafe.Pointer, size uintptr) {
	dstBytes := unsafe.Slice((*byte)(dst), size)
	srcBytes := unsafe.Slice((*byte)(src), size)
	if uintptr(dst) < uintptr(src) || uintptr(dst) >= uintptr(src)+size {
		// No overlap or dst is entirely before src: copy forward
		for i := uintptr(0); i < size; i++ {
			dstBytes[i] = srcBytes[i]
		}
	} else {
		// Overlapping regions, copy backward
		for i := size; i > 0; i-- {
			dstBytes[i-1] = srcBytes[i-1]
		}
	}
}

func unsafe_New(typ reflect.Type) unsafe.Pointer {
	return unsafe.Pointer(&make([]byte, typ.Size()+uintptr(typ.Align()))[0])
}
