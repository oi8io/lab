package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestExamplePrint(t *testing.T) {
	src := `
package main
func main() {
	println("Hello, IoT!")
}
`
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	// Print the AST.
	ast.Print(fset, f)
}

func TestAba(t *testing.T) {
	a := plus()
	a()
	a()
	a()
	fmt.Println(a())
}


func TestAddElement(t *testing.T) {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
func TestFib(t *testing.T) {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}