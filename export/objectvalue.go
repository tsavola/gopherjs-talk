package main // OMIT

type ValueThing struct {
	Foo int
	Bar string
	Baz []uint16
}

func NewValueThing() *ValueThing {
	return &ValueThing{} // HL
}

func (t *ValueThing) GetSomething() int {
	return 42
}
