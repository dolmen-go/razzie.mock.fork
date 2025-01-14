## TLDR

```go
import (
	"testing"
	"github.com/razzie/mock"
)

type Adder interface {
	Add(a, b int) int
}

func TestSomething(t *testing.T) {
	adder, mck := mock.Mock[Adder]() // <---
	mck.On("Add", 1, 2).Return(3)
	res := adder.Add(1, 2)
	if res != 3 {
		t.Errorf("expected 3, got %d", res)
	}
	mck.AssertCalled(t, "Add", 1, 2)
}
```

## What does this package do?

To properly understand the purpose of this package, you must be familiar with [mocking](https://en.wikipedia.org/wiki/Mock_object) and [testify mock](https://pkg.go.dev/github.com/stretchr/testify/mock).

Let's say that in a hypothetical test scenario you need to mock the following `Adder` interface:

```go
type Adder interface {
	Add(a, b int) int
}
```

The usual way to do it would be writing some `MockAdder` boilerplate code very similar to this:

```go
import (
	"testing"
	"github.com/stretchr/testify/mock"
)

type MockAdder struct {
	mock.Mock
}

func (adder *MockAdder) Add(a, b) int {
	ret := adder.Called(a, b)
	return ret.Int(0)
}

func TestSomething(t *testing.T) {
	adder := new(MockAdder)
	adder.On("Add", mock.Anything, mock.Anything).Return(0)
	something(adder)
	adder.AssertCalled(t, "Add", mock.Anything, mock.Anything)
}
```

Maybe it doesn't look like a big deal, but with many large interfaces the work adds up.
Of course you could just run [mockery](https://github.com/vektra/mockery) to generate it,
and if you don't mind occasionally running an extra tool to generate your mocks, then by all means go ahead and use it.

With `github.com/razzie/mock` package though, you can shortcut the work:

```go
import (
	"testing"
	"github.com/razzie/mock"
)

func TestSomething(t *testing.T) {
	var adder, mck = mock.Mock[Adder]() // mck is a testify/mock.Mock object
	mck.On("Add", mock.Anything, mock.Anything).Return(0)
	something(adder)
	mck.AssertCalled(t, "Add", mock.Anything, mock.Anything)
}
```
