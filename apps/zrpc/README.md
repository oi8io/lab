# A GRPC framework like gob

### 实现codec

```golang
type Codec interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(interface{}) error
	Write(header *Header, body interface{}) error
}
```
1. 首先实现io.Closer 的Close方法，这个由TCP conn对象实现。
2. ReadHeader 读取header，把响应报文的数据读取到header中
3. ReadBody读取body，将body的数据读入到传入的对象中
4. Write 写入，将header以及body的内容写入到buf中，而buf则是由conn创建的一个Writer，向远程写入报文。
### 读写操作
1. 读取，先从conn中读取字节流，然decode
2. 写入，将字节进行encode，然后在flush到buf中，发送到远端
```
                   ┌──────────┐           ┌──────────┐         ┌──────────┐
                   │   read   │──────────▶│   conn   ├────────▶│ decoder  │
                   └──────────┘           └──────────┘         └──────────┘
                                                ▲                          
                                                │                          
┌──────────┐        ┌──────────┐          ┌──────────┐                     
│  write   │───────▶│ encoder  │─────────▶│   buf    │                     
└──────────┘        └──────────┘          └──────────┘                     
```



### server端
1. 监听端口，Accept方法阻塞，然后跟进当前参数配置进行codec类型选择
2. 监听请求，如果有请求，则进行解析，并用WaitGroup进行+1进行处理
3. 执行handle

### RPC Call设计
net/rpc一个函数需要能够被远程调用，需要满足5个条件：
* 拥有该方法类型是导出的
* 方法是导出的
* 该方法具有两个参数，都是导出类型或者是内置类型
* 该方法第二个参数是一个指针
* 该方法有一个返回参数error
```go
func (t *T)MethodName (argType T1,replyType *T2) errr
```


### 怎么通过字符串找到执行的方法
```go
			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
			if err := client.Call(ctx, "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
```
这里需要用`Foo.Sum`找到Foo对象，执行Sum方法，只能通过反射执行Method的Call方法执行大致步骤如下：
#### 第一步：注册对象
将对象注册，通过反射，找到对象名称，类型，以及暴露方法列表。
```go
	var wg sync.WaitGroup
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		// j 从 1 开始，第 0 个入参是 wg 自己。
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
```
```go
type service struct {
	name   string
	typ    reflect.Type
	rcvr   reflect.Value
	method map[string]*methodType
}
```
```go
	s := new(service)
	s.rcvr = reflect.ValueOf(rcvr)
	s.name = reflect.Indirect(s.rcvr).Type().Name()
	s.typ = reflect.TypeOf(rcvr)
    s.method = make(map[string]*methodType)
	for i := 0; i < s.rcvr.NumMethod(); i++ {
		method := s.typ.Method(i)
		mType := method.Type
		// 格式必须为 Func(req,*reply) error,判断略
		argv, replyv := mType.In(1), mType.In(2)
		s.method[method.Name] = newMethodType(method, argv, replyv)
		log.Printf("rpc server: register %s.%s\n", s.name, method.Name)
	}
```
#### 第二步，解析请求
当接受到请求后，根据字符串找到相应的service
```go
svci, ok := s.serviceMap.Load(serviceName)
```
```go
func (s *service) call(m *methodType, argv, replyv reflect.Value) error {
	atomic.AddUint64(&m.numCalls, 1)
	f := m.method.Func
	values := f.Call([]reflect.Value{s.rcvr, argv, replyv})
	if err := values[0].Interface(); err != nil {
		return err.(error)
	}
	return nil
}
```


### 执行步骤
#### server端
1. 启动注册进程，获取注册地址
2. 启动servers节点，并注册到注册中心
3. 注册对象，监听请求


#### client端

1. 通过注册中心获取server列表
2. 根据模式选择一个server
3. 发送报文
4. 接受响应


#### 远程调用
1. 客户端发送指定编码的header及body
2. 服务端接受到请求后解析，拿到具体的service
3. 执行根据service 执行Call，将结果打包
4. 发送响应



