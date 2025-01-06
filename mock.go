package mock

import (
	"reflect"
	"unsafe"

	"github.com/stretchr/testify/mock"
	"golang.org/x/exp/constraints"
)

func Mock[T any]() (T, *mock.Mock) {
	typ, methods := getInterfaceTypeAndMethods[T]()
	mockTyp := reflect.StructOf([]reflect.StructField{
		{
			Name: typ.Name() + "Methods",
			Type: reflect.TypeFor[[]*methodInfo](),
		},
		{
			Name: typ.Name() + "Mock",
			Type: reflect.TypeFor[*mock.Mock](),
		},
		{
			Name:      typ.Name(),
			Type:      typ,
			Anonymous: true,
		},
	})
	getMock := func(t reflect.Value) *mock.Mock {
		return t.Field(1).Interface().(*mock.Mock)
	}
	methodInfos := make([]*methodInfo, len(methods))
	abiMethods := getAbiMethods(mockTyp)
	for i, method := range methods {
		fn := icall_array[i]
		fnptr := unsafe.Pointer(reflect.ValueOf(fn).Pointer())
		methodInfos[i] = newMethodInfo(mockTyp, method)
		abiMethods[i].Ifn = addReflectOff(fnptr)
	}

	t := reflect.New(mockTyp).Elem()
	t.Field(0).Set(reflect.ValueOf(methodInfos))
	t.Field(1).Set(reflect.ValueOf(new(mock.Mock)))
	return t.Interface().(T), getMock(t)
}

type abiMethod struct {
	Name int32 // name of method
	Mtyp int32 // method type (without receiver)
	Ifn  int32 // fn used in interface call (one-word receiver)
	Tfn  int32 // fn used for normal method call
}

func getAbiMethods(typ reflect.Type) []abiMethod {
	if typ.Kind() != reflect.Struct {
		panic("not struct kind: " + typ.Kind().String())
	}
	type Type struct {
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
		Name   Name    // name is always non-empty
		Typ    *Type   // type of field
		Offset uintptr // byte offset of field
	}
	type StructType struct {
		Type
		PkgPath Name
		Fields  []StructField
	}
	type StructTypeUncommon struct {
		StructType
		u UncommonType
	}
	type emptyInterface struct {
		typ  *Type
		word unsafe.Pointer
	}
	t := (*emptyInterface)(unsafe.Pointer(&typ)).word
	u := &(*StructTypeUncommon)(t).u
	if u.Mcount == 0 {
		return nil
	}
	m := unsafe.Add(unsafe.Pointer(u), u.Moff)
	return ptrToSlice[abiMethod](m, u.Mcount)
}

type methodInfo struct {
	typ        reflect.Type
	name       string
	inTypes    []reflect.Type
	outTypes   []reflect.Type
	inOffsets  []uintptr
	outOffsets []uintptr
	inSize     uintptr
	outSize    uintptr
}

func newMethodInfo(typ reflect.Type, method reflect.Method) *methodInfo {
	inTypes, outTypes := getMethodInOutTypes(method)
	inOffsets, inSize := argsOffsetsAndSizes(inTypes, true)
	outOffsets, outSize := argsOffsetsAndSizes(outTypes, false)
	return &methodInfo{
		typ:        typ,
		name:       method.Name,
		inTypes:    inTypes,
		outTypes:   outTypes,
		inOffsets:  inOffsets,
		outOffsets: outOffsets,
		inSize:     inSize,
		outSize:    outSize,
	}
}

func icall_n(index int, ptr unsafe.Pointer, p unsafe.Pointer) {
	receiver := reflect.NewAt(reflect.TypeFor[[]*methodInfo](), ptr).Elem()
	info := receiver.Interface().([]*methodInfo)[index]
	receiver = reflect.NewAt(info.typ, ptr).Elem()
	mock := receiver.Field(1).Interface().(*mock.Mock)
	var in []any
	if len(info.inTypes) > 0 {
		buf := make([]byte, info.inSize)
		copy(buf, ptrToSlice[byte](p, info.inSize))
		for i, t := range info.inTypes {
			var arg any
			if t.Size() == 0 {
				arg = reflect.New(t).Elem().Interface()
			} else {
				off := info.inOffsets[i]
				arg = reflect.NewAt(t, unsafe.Pointer(&buf[off])).Elem().Interface()
			}
			in = append(in, arg)
		}
	}
	out := mock.MethodCalled(info.name, in...)
	for i, t := range info.outTypes {
		var v reflect.Value
		if out[i] == nil {
			v = reflect.New(t).Elem()
		} else {
			v = reflect.ValueOf(out[i])
		}
		po := unsafe.Pointer(v.UnsafeAddr())
		memcpy(unsafe.Add(p, info.inSize+info.outOffsets[i]), po, t.Size())
	}
}

var icall_array = []interface{}{
	func(p, a unsafe.Pointer) { icall_n(0, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(1, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(2, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(3, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(4, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(5, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(6, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(7, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(8, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(9, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(10, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(11, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(12, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(13, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(14, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(15, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(16, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(17, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(18, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(19, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(20, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(21, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(22, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(23, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(24, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(25, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(26, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(27, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(28, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(29, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(30, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(31, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(32, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(33, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(34, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(35, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(36, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(37, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(38, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(39, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(40, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(41, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(42, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(43, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(44, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(45, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(46, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(47, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(48, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(49, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(50, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(51, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(52, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(53, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(54, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(55, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(56, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(57, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(58, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(59, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(60, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(61, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(62, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(63, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(64, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(65, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(66, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(67, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(68, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(69, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(70, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(71, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(72, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(73, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(74, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(75, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(76, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(77, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(78, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(79, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(80, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(81, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(82, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(83, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(84, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(85, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(86, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(87, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(88, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(89, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(90, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(91, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(92, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(93, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(94, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(95, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(96, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(97, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(98, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(99, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(100, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(101, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(102, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(103, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(104, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(105, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(106, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(107, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(108, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(109, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(110, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(111, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(112, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(113, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(114, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(115, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(116, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(117, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(118, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(119, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(120, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(121, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(122, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(123, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(124, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(125, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(126, p, unsafe.Pointer(&a)) },
	func(p, a unsafe.Pointer) { icall_n(127, p, unsafe.Pointer(&a)) },
}

//go:linkname addReflectOff reflect.addReflectOff
func addReflectOff(ptr unsafe.Pointer) int32

func ptrToSlice[T any, N constraints.Integer](ptr unsafe.Pointer, len N) []T {
	return unsafe.Slice((*T)(ptr), len)
}

func memcpy[N constraints.Integer](dst, src unsafe.Pointer, len N) {
	copy(ptrToSlice[byte](dst, len), ptrToSlice[byte](src, len))
}

func getInterfaceTypeAndMethods[T any]() (reflect.Type, []reflect.Method) {
	typ := reflect.TypeOf((*T)(nil)).Elem()
	if typ.Kind() != reflect.Interface {
		panic("not an interface type: " + typ.Name())
	}
	methods := make([]reflect.Method, typ.NumMethod())
	for i := range methods {
		methods[i] = typ.Method(i)
	}
	return typ, methods
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

const (
	uintptrAligin = unsafe.Sizeof(uintptr(0))
)

func argsOffsetsAndSizes(typ []reflect.Type, offset bool) (offsets []uintptr, totalSize uintptr) {
	if len(typ) == 0 {
		return nil, 0
	}
	offsets = make([]uintptr, len(typ))
	for i, t := range typ {
		a := uintptr(t.Align())
		totalSize = (totalSize + a - 1) &^ (a - 1)
		offsets[i] = totalSize
		totalSize += t.Size()
	}
	if offset {
		totalSize = (totalSize + uintptrAligin - 1) &^ (uintptrAligin - 1)
		if totalSize == 0 {
			totalSize = uintptrAligin
		}
	}
	return offsets, totalSize
}
