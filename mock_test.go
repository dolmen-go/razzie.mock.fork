package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullableOutput(t *testing.T) {
	type Fooer interface {
		Foo() error
	}
	f, m := Mock[Fooer]()

	m.On("Foo").Once().Return(assert.AnError)
	assert.Equal(t, assert.AnError, f.Foo())
	m.AssertCalled(t, "Foo")

	m.On("Foo").Once().Return(nil)
	assert.NoError(t, f.Foo())
}

func TestNullableInput(t *testing.T) {
	type Fooer interface {
		Foo(a any, b *string)
	}
	f, m := Mock[Fooer]()

	m.On("Foo", nil, (*string)(nil)).Return()
	f.Foo(nil, nil)
	m.AssertCalled(t, "Foo", nil, (*string)(nil))
}

func TestInputArgs(t *testing.T) {
	type Adder interface {
		Add(a, b int) int
	}
	a, m := Mock[Adder]()

	m.On("Add", 1, 2).Return(3)
	assert.Equal(t, 3, a.Add(1, 2))
	m.AssertCalled(t, "Add", 1, 2)
}

func TestVariadicArgs(t *testing.T) {
	type Joiner interface {
		Join(s ...string) string
	}
	j, m := Mock[Joiner]()

	m.On("Join", []string{"a", "b"}).Once().Return("ab")
	assert.Equal(t, "ab", j.Join("a", "b"))
	m.AssertCalled(t, "Join", []string{"a", "b"})

	m.On("Join", []string(nil)).Once().Return("")
	assert.Equal(t, "", j.Join())
}
