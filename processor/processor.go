package processor
import (
       "fmt"
       "github.com/pdelewski/instrgen_testapp/component"
)

type Factory interface {

}

type factory struct {
}

func Type() {
  fmt.Println("processor.Type")
}

func NewProcessor() {
   var f Factory
   f = factory{}
   f.Type()
}