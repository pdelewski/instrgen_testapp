package main

import (
	"github.com/pdelewski/instrgen_test_module/module1"
	"github.com/pdelewski/instrgen_testapp/component"
	"processor"
)

//type Factory interface {
//	component.Factory
//}

type basic struct {
}

func (b basic) Type() {
}

func (b basic) Foo() {
}

func test1() {

	b := basic{}
	b.Type()
	b.Foo()
}

func main() {

	module1.Foo()
	component.Component()
//	var b Factory
//	b = basic{}
//	b.Type()
	test1()
}
