package tree

import "strconv"

var ret [][]string

func PrintTree(root *AVLTree) [][]string {
	if root == nil {
		return nil
	}
	h := getHeight(root)
	w := pow(2, h) - 1
	ret = make([][]string, h)
	for k := range ret {
		s := make([]string, w)
		for key := range s {
			s[key] = ""
		}
		ret[k] = s
	}
	helper(root, 0, 0, w)
	return ret
}
func helper(root *AVLTree, level, l, r int) {
	if root != nil {
		mid := l + (r-l)/2
		ret[level][mid] = strconv.Itoa(root.Value)
		helper(root.Left, level+1, l, mid-1)
		helper(root.Right, level+1, mid+1, r)
	}
}

func getHeight(root *AVLTree) int {
	if root == nil {
		return 0
	}
	l := getHeight(root.Left)
	r := getHeight(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
func pow(x, y int) int {
	ret := x
	for i := 1; i < y; i++ {
		ret *= x
	}
	return ret
}
