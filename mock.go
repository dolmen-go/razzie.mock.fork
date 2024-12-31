package mock

import (
	"reflect"

	"github.com/goplus/reflectx"
	"github.com/stretchr/testify/mock"
)

func Mock[T any]() (T, *mock.Mock) {
	methods := getInterfaceMethods[T]()
	mockTyp := reflect.StructOf([]reflect.StructField{
		{
			Name:      "Mock",
			Type:      reflect.TypeFor[mock.Mock](),
			Anonymous: true,
		},
	})
	getMock := func(ptr reflect.Value) *mock.Mock {
		return ptr.Elem().FieldByName("Mock").Addr().Interface().(*mock.Mock)
	}
	mockMethodSet := reflectx.NewMethodSet(mockTyp, 0, len(methods))
	mockMethods := make([]reflectx.Method, 0, len(methods))
	for _, method := range methods {
		outTypes := getMethodOutTypes(method)
		impl := func(in []reflect.Value) []reflect.Value {
			m := getMock(in[0])
			in = in[1:] // skip receiver
			out := m.MethodCalled(method.Name, reflectValuesToInterfaces(in)...)
			return interfacesToReflectValues(out, outTypes)
		}
		mockMethods = append(mockMethods, reflectx.Method{
			Name:    method.Name,
			PkgPath: method.PkgPath,
			Pointer: true,
			Type:    method.Type,
			Func:    impl,
		})
	}
	reflectx.SetMethodSet(mockMethodSet, mockMethods, false)

	t := reflect.New(mockMethodSet)
	return t.Interface().(T), getMock(t)
}

func getInterfaceMethods[T any]() []reflect.Method {
	typ := reflect.TypeOf((*T)(nil)).Elem()
	if typ.Kind() != reflect.Interface {
		panic("not an interface type: " + typ.Name())
	}
	methods := make([]reflect.Method, typ.NumMethod())
	for i := range methods {
		methods[i] = typ.Method(i)
	}
	return methods
}

func getMethodOutTypes(method reflect.Method) []reflect.Type {
	typ := method.Type
	len := typ.NumOut()
	out := make([]reflect.Type, 0, len)
	for i := 0; i < len; i++ {
		out = append(out, typ.Out(i))
	}
	return out
}

func reflectValuesToInterfaces(in []reflect.Value) []any {
	out := make([]any, 0, len(in))
	for _, value := range in {
		out = append(out, value.Interface())
	}
	return out
}

func interfacesToReflectValues(in []any, types []reflect.Type) []reflect.Value {
	out := make([]reflect.Value, 0, len(in))
	for i, intf := range in {
		if intf == nil {
			out = append(out, reflect.New(types[i]).Elem())
			continue
		}
		out = append(out, reflect.ValueOf(intf))
	}
	return out
}
