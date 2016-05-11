package main // OMIT

import "github.com/gopherjs/gopherjs/js" // OMIT

type ActiveThing struct {
	Foo int
	Bar string
	Baz []uint16
}

func NewActiveThing() *js.Object {
	return js.MakeWrapper(&ActiveThing{}) // HL
}

func (t *ActiveThing) GetSomething() int {
	return 42
}
