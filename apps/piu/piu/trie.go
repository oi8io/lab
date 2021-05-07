package piu

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

func (n *node) toString() {
	fmt.Printf("pattern (%s) ,part (%s), isWild (%v) \n", n.pattern, n.part, n.isWild)
	for _, child := range n.children {
		child.toString()
	}
}

// matchChild 完全匹配 或者 支持匹配，则匹配成功
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if part == child.part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 找出匹配成功的所有child
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if part == child.part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//insert
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height { //长度与处理层数相等， 证明已经处理完毕。
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil { // 未匹配成功则创建新的子节点
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1) // 下一层
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1) // 递归 匹配
		if result != nil {
			return result
		}
	}
	return nil
}
