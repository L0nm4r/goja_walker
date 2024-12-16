package gojawalker

import (
	"fmt"
	"reflect"

	"github.com/dop251/goja/ast"
)

type ExampleVisitor struct{}

func (v ExampleVisitor) Enter(n ast.Node) IVisitor {
	nodeType := reflect.TypeOf(n)
	fmt.Printf("enter node: \n---- ---- ----\n\t%s,\n\t[type: %s]\n", n, nodeType)
	return v
}

func (v ExampleVisitor) Exit(n ast.Node) {
	nodeType := reflect.TypeOf(n)
	fmt.Printf("leaving node: \n---- ---- ----\n\t%s,\n\t[type: %s]\n", n, nodeType)
}
