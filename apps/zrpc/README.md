# A GRPC framework like gob

### 实现codec

```golang
type Codec interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(interface{}) error
	Writer(header *Header, body interface{}) error
}
```
1. 首先实现io.Closer 的Close方法，这个由TCP conn对象实现。
2. ReadHeader 读取header，把响应报文的数据读取到header中
3. ReadBody读取body，将body的数据读入到传入的对象中
4. Writer 写入，将header以及body的内容写入到buf中，而buf则是由conn创建的一个Writer，向远程写入报文。
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

