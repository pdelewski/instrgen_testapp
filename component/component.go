package component

import "fmt"

type Factory interface {
  Type()
}

func Component() {
  fmt.Println("Component")
}