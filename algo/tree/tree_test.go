package tree

import (
	"testing"
)

func TestEcho(t *testing.T) {
	Echo()
}

func TestNewTree(t *testing.T) {
	dir := "/Users/ocean/develop/go/o2ln"
	tree := NewTree(nil, dir)
	//fmt.Println(tree)
	tree.midEcho()
}
