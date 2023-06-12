package main

import (
	"github.com/pdelewski/instrgen_test_module/module1"
	"github.com/pdelewski/instrgen_testapp/component"
	"github.com/pdelewski/instrgen_testapp/processor"
)

type Factory interface {
  component.Factory2
}

type factory2 struct {
}

type factory struct {
}

func (f factory2) Type2() {

}

func (f factory) Type() {
}

func test1() {
  var f Factory
  f = factory2{}
  f.Type2()
}



func main() {

	module1.Foo()
	component.Component()

	var f component.Factory
	f = factory{}
	f.Type()
	test1()

	processor.NewProcessor()
}
