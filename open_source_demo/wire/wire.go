//+build wireinject

package wire

import (
	"fmt"
	wire "github.com/google/wire"
)

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
	foo *Foo
}

func NewBar(foo *Foo) *Bar {
	return &Bar{
		foo: foo,
	}
}

func (p *Bar) Test() {
	fmt.Println("hello")
}

type Instance struct {
	Foo *Foo
	Bar *Bar
}

var SuperSet = wire.NewSet(NewFoo, NewBar)

func InitializeAllInstance() *Instance {
	wire.Build(SuperSet, Instance{})
	return &Instance{}
}
