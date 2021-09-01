package huffman

import (
	"oi.io/algo/tree"
	"sort"
)

type Huffman struct {
	data       string
	dataStatic map[int32]int
	arr        []int
	tree       *tree.AVLTree
	code       []byte
}

func NewHuffman(str string) *Huffman {
	var h = new(Huffman)
	h.data = str
	h.dataStatic = make(map[int32]int)
	for _, o := range str {
		if _, ok := h.dataStatic[o]; ok {
			h.dataStatic[o] = h.dataStatic[o] + 1
		} else {
			h.dataStatic[o] = 1
		}
	}
	var arrOfTree PairList
	for o, y := range h.dataStatic {
		h.arr = append(h.arr, y)
		t := tree.AVLTree{}
		t.Value = y
		t.PayLoad = o
		t.Height = 0
		arrOfTree = append(arrOfTree, t)
	}
	sort.Sort(arrOfTree)
	//fmt.Println(arrOfTree)
	tree := GetMinTree(arrOfTree)
	h.tree = &tree
	return h
}

type PairList []tree.AVLTree

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func GetHeight(a, b int) int {
	if a > b {
		return a + 1
	} else {
		return b + 1
	}
}

func SliceInsert(s []tree.AVLTree, index int, value tree.AVLTree) []tree.AVLTree {
	rear := append([]tree.AVLTree{}, s[index:]...)
	return append(append(s[:index], value), rear...)
}

//GetMinTree
func GetMinTree(arr []tree.AVLTree) tree.AVLTree {
	for {
		if len(arr) < 2 {
			//fmt.Println("arr=>", arr)
			return arr[0]
		}
		sum := arr[0].Value + arr[1].Value
		//fmt.Println(arr[0].Value, "<===", sum, "===>", arr[1].Value)
		element := tree.AVLTreeElement{Value: sum, Height: GetHeight(arr[0].GetHeight(), arr[1].GetHeight())}
		t := tree.AVLTree{AVLTreeElement: element, Left: &arr[0], Right: &arr[1]}
		arr = arr[2:]
		//fmt.Println(arr)
		if len(arr) == 0 {
			//fmt.Println("t=>>>", t.Right)
			return t
		}
		point := len(arr)
		for k, v := range arr {
			if sum < v.Value {
				//fmt.Println("sum=>", sum, v.Value)
				point = k
				break
			}
		}
		arr = SliceInsert(arr, point, t)
		//fmt.Println(arr)
	}
}

func getTwoMinTree(bt *tree.AVLTree, arr []int) (a, b int, ret []int) {
	ak, bk := -1, -1
	for k, v := range arr {
		if ak == -1 {
			ak = k
			a = v
			continue
		}
		if bk == -1 {
			bk = k
			b = v
			continue
		}
		if v < a {
			b = a
			a = v
			continue
		}
		if v < b {
			b = v
			continue
		}
	}
	if ak < bk {
		arr = append(arr[:ak-1], arr[ak+1:]...)
		arr = append(arr[:bk], arr[bk:]...)
	} else {
		arr = append(arr[:bk-1], arr[ak+1:]...)
		arr = append(arr[:ak], arr[ak:]...)
	}

	ret = arr
	return
}

func GetTwoMinNum(arr []*tree.AVLTree) (ak, bk int) {
	ak, bk = -1, -1
	a, b := 0, 0
	for k, r := range arr {
		v := r.Value
		if ak == -1 {
			ak = k
			a = v
			continue
		}
		if bk == -1 {
			bk = k
			b = v
			continue
		}
		if v < a {
			b = a
			a = v
			continue
		}
		if v < b {
			b = v
			continue
		}
	}
	if ak > bk {
		return bk, ak
	}
	return
}
