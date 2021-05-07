package piu

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Params     map[string]string
	Path       string
	Method     string
	StatusCode int
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
	return &Context{Writer: writer, Request: request, Path: request.RequestURI, Method: request.Method}
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

func (c *Context) Html(statusCode int, html string) {
	c.Status(statusCode)
	c.Writer.Write([]byte(html))
}

func (c *Context) Json(statusCode int, data interface{}) {
	c.Status(statusCode)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
