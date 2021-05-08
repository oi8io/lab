package piu

import (
	"fmt"
	"strings"
)

type iRouter interface {
	AddRoute(method, pattern string, handlerFunc HandlerFunc)
	GetRoute(method, pattern string) (handlerFunc HandlerFunc)
	GetRouteKey(method, pattern string) string
}

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: map[string]HandlerFunc{}, roots: map[string]*node{}}
}

func (r *router) GetRouteKey(method, pattern string) string {
	return fmt.Sprintf("%s_%s", method, pattern)
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handlerFunc HandlerFunc) {
	parts := parsePattern(pattern)

	key := r.GetRouteKey(method, pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handlerFunc

	r.handlers[key] = handlerFunc
}

func (r *router) getRoute(method, path string) (node *node, params map[string]string) {
	searchParts := parsePattern(path)
	params = make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) Handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path) //匹配路由
	c.Params = params
	if node == nil { //  找不到路由，将404追加到里面
		c.handlers = append(c.handlers,NotFound)
	}else {
		key := r.GetRouteKey(c.Method, node.pattern)
		handler := r.handlers[key]
		// 找到路由，当前路由方法追加到里面
		c.handlers = append(c.handlers, handler)
	}
	// 开始执行
	c.Next()
}
