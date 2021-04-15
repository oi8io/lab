package errcode

type Error struct {
	msg  string
	code int
}

func NewError(code int, msg string) *Error {
	return &Error{code: 1000 + code, msg: msg}
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() int {
	return e.code
}

var NotModified = NewError(1, "Not Modified")
var NotFound = NewError(2, "not found")
var ParamError = NewError(3, "Param Error")
