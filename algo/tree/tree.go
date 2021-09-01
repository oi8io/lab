package tree

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//Element 
type Element struct {
	Name  string
	IsDir bool
	Size  int64
}

func NewElement(name string, isDir bool, size int64) *Element {
	return &Element{Name: name, IsDir: isDir, Size: size}
}

type Tree struct {
	Data        interface{}
	FirstChild  *Tree
	NextSibling *Tree
}

//ADTofTree 树的结构
type ADTofTree struct {
	Data        interface{} // 节点数据
	FirstChild  *ADTofTree  // 第一个子节点
	NextSibling *ADTofTree  // 下一个兄弟节点
}

func NewTree(parent *Tree, dir string) *Tree {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	if parent == nil {
		parent = new(Tree)
		data := NewElement("root", true, 0)
		parent.Data = data
	}
	//fmt.Println("dir", dir)
	before := new(Tree)
	for _, f := range files {
		data := NewElement(f.Name(), f.IsDir(), f.Size())
		child := new(Tree)
		child.Data = data
		if f.IsDir() {
			child = NewTree(child, dir+"/"+f.Name())
		}
		if parent.FirstChild == nil {
			//fmt.Println(parent.Data.name + " add firstChild " + f.Name())
			parent.FirstChild = child
		} else {
			//fmt.Println(before.Data.name + " add brother " + f.Name())
			before.NextSibling = child
		}
		before = child
	}
	return parent
}

func (t *Tree) preEcho(sub int) {
	if t == nil {
		return
	}
	pre := strings.Repeat("  ", sub) + "|->"
	fmt.Println(pre + t.Data.(Element).Name)
	if t.Data.(Element).IsDir {
		t.FirstChild.preEcho(sub + 1)
		//t.preEcho(1)
	} else {

	}
	if t.NextSibling != nil {
		tmp := t.NextSibling
		tmp.preEcho(sub)
	}
}

func (t *Tree) midEcho() {

}
func (t *Tree) lasEcho(sub int) int64 {
	if t == nil {
		return 0
	}
	sum := t.Data.(Element).Size
	if t.NextSibling != nil {
		tmp := t.NextSibling
		sum += tmp.lasEcho(sub)
	}

	pre := strings.Repeat("  ", sub) + "|->"
	if t.Data.(Element).IsDir {
		sum += t.FirstChild.lasEcho(sub + 1)

		//t.preEcho(1)
	} else {
		//fmt.Println("->",t.Data.size)
	}
	fmt.Println(pre + t.Data.(Element).Name + " " + strconv.FormatInt(sum, 10))
	return sum
}

func Echo() {
	fmt.Println("echo")
}
