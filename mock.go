package mock

import (
	"reflect"

	"github.com/goplus/reflectx"
	"github.com/stretchr/testify/mock"
)

func Mock[T any]() (T, *mock.Mock) {
	_, methods := getInterfaceTypeAndMethods[T]()
	mockTyp := reflect.StructOf([]reflect.StructField{
		{
			Name:      "Mock",
			Type:      reflect.TypeFor[mock.Mock](),
			Anonymous: true,
		},
	})
	mockMethodSet := reflectx.NewMethodSet(mockTyp, 0, len(methods))
	mockMethods := make([]reflectx.Method, 0, len(methods))
	for _, method := range methods {
		impl := func(in []reflect.Value) []reflect.Value {
			m := in[0].Elem().FieldByName("Mock").Addr().Interface().(*mock.Mock)
			in = in[1:] // skip receiver
			out := m.MethodCalled(method.Name, reflectValuesToInterfaces(in)...)
			return interfacesToReflectValues(out)
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
	return t.Interface().(T), t.Elem().FieldByName("Mock").Addr().Interface().(*mock.Mock)
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

func reflectValuesToInterfaces(in []reflect.Value) []any {
	out := make([]any, 0, len(in))
	for _, value := range in {
		out = append(out, value.Interface())
	}
	return out
}

func interfacesToReflectValues(in []any) []reflect.Value {
	out := make([]reflect.Value, 0, len(in))
	for _, intf := range in {
		out = append(out, reflect.ValueOf(intf))
	}
	return out
}
