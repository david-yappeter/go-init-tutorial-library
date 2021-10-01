package goinittutoriallibrary

import (
	"fmt"

	"github.com/david-yappeter/go-init-tutorial-library/nested"
)

var Var = 1

func init() {
	_ = nested.A
	fmt.Println("library root package")
}
