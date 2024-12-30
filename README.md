```go
import (
	"testing"
	"github.com/razzie/mock"
)

type Fooer interface {
	Foo()
}

func TestFooer(t *testing.T) {
	f, m := mock.Mock[Fooer]()
	m.On("Foo").Return()
	f.Foo()
	m.AssertCalled(t, "Foo")
}
```