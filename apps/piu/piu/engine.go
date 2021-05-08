package piu

import (
	"html/template"
	"net/http"
	"strings"
)

var _ http.Handler = NewEngine()

type HandlerFunc func(c *Context)

type H map[string]interface{}

//Engine 一个http.Handler，控制程序的运行
type Engine struct {
	*RouterGroup                         // 继承所有RouterGroup的功能
	route           *router              // 路由对象
	groups          []*RouterGroup       // 路由分组
	template        *template.Template   //模板对象
	templatePattern string               //模板文件位置（略）
	funcMap         template.FuncMap     //模板函数合集
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

// 请求入口
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	// 找到当前中间件
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	// 新建一个context控制该请求
	ctx := NewContext(w, r)
	ctx.engine = e
	ctx.handlers = middlewares //指定中间件
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
