package simple

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}

type FooBar struct {
	*Foo
	*Bar
}

func NewFooBar(foo *Foo, bar *Bar) *FooBar {
	return &FooBar{
		Foo: foo,
		Bar: bar,
	}
}
