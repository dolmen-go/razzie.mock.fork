```go
package main

import "github.com/razzie/mock"

type Fooer interface {
	Foo()
}

func main() {
	f, m := mock.Mock[Fooer]()
	m.On("Foo").Once().Return()
	f.Foo()
	m.AssertCalled(t, "Foo")
}
```