package mock

import (
	"reflect"

	"github.com/stretchr/testify/mock"
)

// Anything is used when the argument being tested
// shouldn't be taken into consideration.
const Anything = mock.Anything

// Mock takes a T interface type and returns a runtime generated mock implementation for it
// plus a testify mock.Mock object to control its behavior
func Mock[T any]() (T, *mock.Mock) {
	typ := reflect.TypeOf((*T)(nil)).Elem()
	if typ.Kind() != reflect.Interface {
		panic("not an interface type: " + typ.Name())
	}

	mockTyp := reflect.StructOf([]reflect.StructField{
		{
			Name: "IcallBase__" + typ.Name(),
			Type: reflect.TypeFor[icallBase](),
		},
		{
			Name: "Mock__" + typ.Name(),
			Type: reflect.TypeFor[*mock.Mock](),
		},
		{
			Name:      typ.Name(),
			Type:      typ,
			Anonymous: true,
		},
	})

	m := new(mock.Mock)
	methodInfos := make(icallBase, typ.NumMethod())
	abiMethods := getAbiMethods(mockTyp)
	for i := range methodInfos {
		methodInfos[i] = newMethodInfo(mockTyp, typ.Method(i))
		abiMethods[i].setFn(icall_array[i])
	}

	t := reflect.New(mockTyp).Elem()
	t.Field(0).Set(reflect.ValueOf(methodInfos))
	t.Field(1).Set(reflect.ValueOf(m))
	return t.Interface().(T), m
}
