package piu

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	//r.addRoute("GET", "/hi/:name", nil)
	//r.addRoute("GET", "/assets/*filepath", nil)
	for s, n := range r.roots {
		fmt.Println(s)
		n.toString()
	}
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/piu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "piu" {
		t.Fatal("name should be equal to 'piu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want *router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_AddRoute(t *testing.T) {
	type fields struct {
		roots    map[string]*node
		handlers map[string]HandlerFunc
	}
	type args struct {
		method      string
		pattern     string
		handlerFunc HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			fmt.Println(r)
		})
	}
}

func TestRouter_GetRoute(t *testing.T) {
	type fields struct {
		roots    map[string]*node
		handlers map[string]HandlerFunc
	}
	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantNode   *node
		wantParams map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			gotNode, gotParams := r.getRoute(tt.args.method, tt.args.path)
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("getRoute() gotNode = %v, want %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotParams, tt.wantParams) {
				t.Errorf("getRoute() gotParams = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}

func TestRouter_GetRouteKey(t *testing.T) {
	type fields struct {
		roots    map[string]*node
		handlers map[string]HandlerFunc
	}
	type args struct {
		method  string
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			if got := r.GetRouteKey(tt.args.method, tt.args.pattern); got != tt.want {
				t.Errorf("GetRouteKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_Handle(t *testing.T) {
	type fields struct {
		roots    map[string]*node
		handlers map[string]HandlerFunc
	}
	type args struct {
		c *Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			fmt.Println(r)
		})
	}
}

func Test_parsePattern(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
