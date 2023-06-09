package processor
import (
       "fmt"
       "github.com/pdelewski/instrgen_testapp/component"
)

type Factory interface {
  component.Factory
}

type factory struct {
}

func(f factory) Type() {
  fmt.Println("processor.Type")
}

func NewProcessor() {
   var f Factory
   f = factory{}
   f.Type()
}