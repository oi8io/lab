package zrpc

import (
	"go/ast"
	"log"
	"reflect"
	"sync/atomic"
)

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
	numCalls  uint64
}

func newMethodType(method reflect.Method, argType reflect.Type, replyType reflect.Type) *methodType {
	return &methodType{method: method, ArgType: argType, ReplyType: replyType}
}

func (m *methodType) NumCalls() uint64 {
	return atomic.LoadUint64(&m.numCalls)
}

func (m *methodType) newArgv() reflect.Value {
	var argv reflect.Value
	if m.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(m.ArgType.Elem())
	} else {
		argv = reflect.New(m.ArgType).Elem()
	}
	return argv
}

func (m *methodType) newReplyv() reflect.Value {
	var reply = reflect.New(m.ReplyType.Elem())
	switch m.ReplyType.Elem().Kind() {
	case reflect.Map:
		reply.Elem().Set(reflect.MakeMap(m.ReplyType.Elem()))
	case reflect.Slice:
		reply.Elem().Set(reflect.MakeSlice(m.ReplyType.Elem(), 0, 0))
	}
	return reply
}

type service struct {
	name   string
	typ    reflect.Type
	rcvr   reflect.Value
	method map[string]*methodType
}

func newService(rcvr interface{}) *service {
	s := new(service)
	s.rcvr = reflect.ValueOf(rcvr)
	s.name = reflect.Indirect(s.rcvr).Type().Name()
	s.typ = reflect.TypeOf(rcvr)
	if !ast.IsExported(s.name) {
		log.Fatalf("rpc server: %s is not a valid service name", s.name)
	}
	s.registerMethod()
	return s
}
func (s *service) registerMethod() {
	s.method = make(map[string]*methodType)
	for i := 0; i < s.rcvr.NumMethod(); i++ {
		method := s.typ.Method(i)
		mType := method.Type
		if mType.NumIn() != 3 || mType.NumOut() != 1 {
			continue
		}
		if mType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
			continue
		}
		argv, replyv := mType.In(1), mType.In(2)

		if !isExportedOrBuiltinType(argv) || !isExportedOrBuiltinType(replyv) {
			continue
		}

		s.method[method.Name] = newMethodType(method, argv, replyv)
		log.Printf("rpc server: register %s.%s\n", s.name, method.Name)
	}
}

func (s *service) call(m *methodType, argv, replyv reflect.Value) error {
	atomic.AddUint64(&m.numCalls, 1)
	f := m.method.Func
	values := f.Call([]reflect.Value{s.rcvr, argv, replyv})
	if err := values[0].Interface(); err != nil {
		return err.(error)
	}
	return nil
}

func isExportedOrBuiltinType(t reflect.Type) bool {
	return ast.IsExported(t.Name()) || t.PkgPath() == ""
}
