package piu

import (
	"net/http"
)

var _ http.Handler = NewEngine()

type HandlerFunc func(c *Context)

type H map[string]interface{}

type Engine struct {
	*RouterGroup
	route  *router
	groups []*RouterGroup
}

func (e *Engine) Use(handlerFunc HandlerFunc) {
	e.route.addRoute(http.MethodGet, "pattern", handlerFunc)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func NewEngine() *Engine {
	engine := &Engine{route: newRouter()}
	engine.RouterGroup = NewRouterGroup(engine)
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	e.route.Handle(ctx)
}
