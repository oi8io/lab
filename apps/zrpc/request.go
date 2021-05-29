package zrpc

import (
	"oi.io/apps/zrpc/codec"
	"reflect"
)

type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
	mtype        *methodType
	svc          *service
}
