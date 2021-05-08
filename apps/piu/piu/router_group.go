package piu

import (
	"net/http"
	"path"
)

type iRouterGroup interface {
	Get(pattern string, handlerFunc HandlerFunc)
	Put(pattern string, handlerFunc HandlerFunc)
	Post(pattern string, handlerFunc HandlerFunc)
	Delete(pattern string, handlerFunc HandlerFunc)
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	group := NewRouterGroup(r.engine)
	group.prefix = r.prefix + prefix
	group.parent = r
	group.engine.groups = append(group.engine.groups, group)
	return group
}

func (r *RouterGroup) Use(middleware ...HandlerFunc) {
	r.middlewares = append(r.middlewares, middleware...)
}

func (r *RouterGroup) AddRouter(method, pattern string, handlerFunc HandlerFunc) {
	r.engine.route.addRoute(method, r.prefix+pattern, handlerFunc)
}

func (r *RouterGroup) Get(pattern string, handlerFunc HandlerFunc) {
	r.AddRouter(http.MethodGet, pattern, handlerFunc)
}

func (r *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(r.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

func (r *RouterGroup) Static(relativePath string, root string) {
	handler := r.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	r.Get(urlPattern, handler)
}

func (r *RouterGroup) Put(pattern string, handlerFunc HandlerFunc) {
	r.AddRouter(http.MethodPut, pattern, handlerFunc)
}

func (r *RouterGroup) Post(pattern string, handlerFunc HandlerFunc) {
	r.AddRouter(http.MethodPost, pattern, handlerFunc)
}

func (r *RouterGroup) Delete(pattern string, handlerFunc HandlerFunc) {
	r.AddRouter(http.MethodDelete, pattern, handlerFunc)
}

func NewRouterGroup(engine *Engine) *RouterGroup {
	return &RouterGroup{engine: engine}
}
