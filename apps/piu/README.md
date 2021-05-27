#  web framework like gin

## 结构介绍
### 核心结构：
1. `Engine内核引擎` 实现http.Handler接口
2. `Router路由` 处理GET,POST，PUT,DELETE）
3. `Context` 处理上下文，请求及响应载体
4. `Trie前缀（字典）树` 路由规则匹配
### Engine 介绍
Engine 除了实现http.Handler 接口外，还需要具有所有router所具有的功能，其中还额外包括template处理，服务启动，中间件等功能。
```golang
//Engine 一个http.Handler，控制程序的运行
type Engine struct {
    *RouterGroup   // 继承所有RouterGroup的功能
    route           *router  // 路由对象
    groups          []*RouterGroup // 路由分组
    template        *template.Template //模板对象
    templatePattern string //模板文件位置（略）
    funcMap         template.FuncMap //模板函数合集
}
```
#### http.Handler的实现
```golang
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
```
### Context介绍
### 结构
```golang
//Context 跟随整个请求生命周期，处理请求，每个请求一个context
type Context struct {
	Write     http.ResponseWriter   // writer
	Request    *http.Request         // request
	Params     map[string]string     // url参数 如/hello/:name 访问 /hello/piu 则params[name]=piu
	Path       string                // URI /hello/piu
	Method     string                // GET POST PUT DELETE 其他暂时不支持
	StatusCode int                   // 响应状态码 200 401 etc.
	handlers   []HandlerFunc         // 中间件及最终执行的handler
	index      int                   // index，当前执行到第几个handler了
	engine     *Engine               // 引擎，主要获取全局template对象
}
```
#### 中间件执行过程
```golang
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
    // 开始执行 进入套娃模式
	c.Next()
}
```
```golang
func (c *Context) Next() {
	c.index++ 
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c) //分别执行
	}
}

```
写中间件代码如下：
```golang
func Mid1() HandlerFunc {
	name := "Mid1"
	return func(c *Context) {
		fmt.Println("start run", name)
		c.Next()
		fmt.Println("end   run", name)
	}
}
// 并使用
engine.Use(piu.Mid1())
engine.Use(piu.Mid2())
engine.Use(piu.Mid3())
```
执行顺序如下：
```text
start run Mid1
start run Mid2
start run Mid3
execting
end   run Mid3
end   run Mid2
end   run Mid1
```


### todo