package piu

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"net/http"
	"time"
)

//Context 跟随整个请求生命周期，处理请求，每个请求一个context
type Context struct {
	Writer     http.ResponseWriter   // writer
	Request    *http.Request         // request
	Params     map[string]string     // url参数 如/hello/:name 访问 /hello/piu 则params[name]=piu
	Path       string                // URI /hello/piu
	Method     string                // GET POST PUT DELETE 其他暂时不支持
	StatusCode int                   // 响应状态码 200 401 etc.
	handlers   []HandlerFunc         // 中间件及最终执行的handler
	index      int                   // index，当前执行到第几个handler了
	engine     *Engine               // 引擎，主要获取全局template对象
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Context) Done() <-chan struct{} {
	c2 := make(chan struct{})
	return c2
}

func (c *Context) Err() error {
	return errors.New("error")
}

func (c *Context) Value(key interface{}) interface{} {
	return "value x"
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{Writer: writer, Request: request, Path: request.RequestURI, Method: request.Method, index: -1}
}

func (c *Context) AddFuncMap(name string, function interface{}) {
	c.engine.funcMap[name] = function
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Form(key string) string {
	return c.Request.Form.Get(key)
}

func (c *Context) GetBody() ([]byte, error) {
	return io.ReadAll(c.Request.Body)
}

func (c *Context) Abort(statusCode int) {
	c.Writer.WriteHeader(statusCode)
}

func (c *Context) Status(statusCode int) {
	c.StatusCode = statusCode
	c.Writer.WriteHeader(statusCode)
}

func (c *Context) SetHeader(key, val string) {
	c.Writer.Header().Set(key, val)
}

func (c *Context) Data(statusCode int, data []byte) {
	c.Status(statusCode)
	c.Writer.Write(data)
}

func (c *Context) String(statusCode int, content string) {
	c.Status(statusCode)
	c.Writer.Write([]byte(content))
}

func (c *Context) Fail(statusCode int, content string) {
	c.Status(statusCode)
	c.Writer.Write([]byte(content))
}

func (c *Context) getTemplate() *template.Template {
	return c.engine.getTemplate()
}

func (c *Context) HTML(statusCode int, name string, data interface{}) {
	c.Status(statusCode)
	c.SetHeader("Content-Type", "text/html")
	if err := c.getTemplate().ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(http.StatusInternalServerError, err.Error())
	}
}

func (c *Context) Json(statusCode int, data interface{}) {
	c.Status(statusCode)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}
