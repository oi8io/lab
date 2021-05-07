package piu

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type iEngine interface {
	Get(pattern string, handlerFunc HandlerFunc)
	Put(pattern string, handlerFunc HandlerFunc)
	Post(pattern string, handlerFunc HandlerFunc)
	Delete(pattern string, handlerFunc HandlerFunc)
	Use(handlerFunc HandlerFunc)
	Run(addr string) error
}

type Engine struct {
	route *router
}

func (e *Engine) Use(handlerFunc HandlerFunc) {
	e.route.addRoute(http.MethodGet, "pattern", handlerFunc)
}

func (e *Engine) Get(pattern string, handlerFunc HandlerFunc) {
	e.route.addRoute(http.MethodGet, pattern, handlerFunc)
}

func (e *Engine) Put(pattern string, handlerFunc HandlerFunc) {
	e.route.addRoute(http.MethodPut, pattern, handlerFunc)
}

func (e *Engine) Post(pattern string, handlerFunc HandlerFunc) {
	e.route.addRoute(http.MethodPost, pattern, handlerFunc)
}

func (e *Engine) Delete(pattern string, handlerFunc HandlerFunc) {
	e.route.addRoute(http.MethodDelete, pattern, handlerFunc)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func NewEngine() *Engine {
	return &Engine{route: newRouter()}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	e.route.Handle(ctx)
}
