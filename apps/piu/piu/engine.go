package piu

import (
	"html/template"
	"net/http"
	"strings"
)

var _ http.Handler = NewEngine()

type HandlerFunc func(c *Context)

type H map[string]interface{}

type Engine struct {
	*RouterGroup
	route           *router
	groups          []*RouterGroup
	template        *template.Template
	templatePattern string
	funcMap         template.FuncMap
}

func (e *Engine) Use(handlerFunc ...HandlerFunc) {
	e.middlewares = append(e.middlewares, handlerFunc...)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func NewEngine() *Engine {
	engine := &Engine{route: newRouter()}
	engine.RouterGroup = NewRouterGroup(engine)
	engine.groups = []*RouterGroup{engine.RouterGroup}
	engine.funcMap = make(template.FuncMap)
	return engine
}

func NotFound(c *Context) {
	c.Status(http.StatusNotFound)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc

	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	ctx := NewContext(w, r)
	ctx.engine = e
	ctx.handlers = middlewares
	e.route.Handle(ctx)
}

func (e *Engine) getTemplate() *template.Template {
	return e.template
}

func (e *Engine) AddFuncMap(name string, function interface{}) {
	e.funcMap[name] = function
}

func (e *Engine) LoadHTMLGlob(pattern string) { //todo 最后执行 funcMap
	e.templatePattern = pattern
	e.template = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(e.templatePattern))
}
