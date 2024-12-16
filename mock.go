package mock

import (
	"reflect"

	"github.com/stretchr/testify/mock"
)

func Mock[T any]() (T, *mock.Mock) {
	/*_, methods := getInterfaceTypeAndMethods[T]()
	fields := make([]reflect.StructField, 0, len(methods))
	for _, method := range methods {
		fields = append(fields, reflect.StructField{
			Name:    method.Name,
			PkgPath: method.PkgPath,
			Type:    method.Type,
		})
	}
	mTyp := reflect.StructOf(fields)
	t := reflect.New(mTyp).Elem()
	m := new(mock.Mock)
	for i, method := range methods {
		impl := func(in []reflect.Value) []reflect.Value {
			in = in[1:] // ignore receiver
			out := m.MethodCalled(method.Name, reflectValuesToInterfaces(in)...)
			return interfacesToReflectValues(out)
		}
		t.Field(i).Set(reflect.MakeFunc(method.Type, impl))
	}
	return t.Addr().Interface().(T), m*/

	typ, methods := getInterfaceTypeAndMethods[T]()
	mTyp := reflect.StructOf([]reflect.StructField{
		{
			Name:      typ.Name(),
			Type:      typ,
			Anonymous: true,
		},
	})
	t := reflect.New(mTyp).Elem()
	m := new(mock.Mock)
	for _, method := range methods {
		impl := func(in []reflect.Value) []reflect.Value {
			in = in[1:] // ignore receiver
			out := m.MethodCalled(method.Name, reflectValuesToInterfaces(in)...)
			return interfacesToReflectValues(out)
		}
		t.MethodByName(method.Name).Set(reflect.MakeFunc(method.Type, impl))
	}
	return t.Addr().Interface().(T), m
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
