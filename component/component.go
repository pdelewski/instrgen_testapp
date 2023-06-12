package component

import (
	"fmt"
)

type Factory interface {
	Type()
}

type Factory2 interface {
        Type2()
}

type factory struct {
}

func (f factory) Type() {

	fmt.Println("component.Factory.Type()")
}

func Component() {

	fmt.Println("Component")
	var f Factory
	f = factory{}
	f.Type()
}
