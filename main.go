package main

import "github.com/pdelewski/instrgen_test_module/module1"
import "github.com/pdelewski/instrgen_testapp/component"

type Factory interface {
  component.Factory
}

type basic struct {
}

func (b basic) Type() {
}

func main() {
  module1.Foo()
  component.Component()
  var b Factory
  b = basic{}
  b.Type()
}