package utils

import "fmt"

type UserError struct {
	Code int
	Msg  string
}

func (u UserError) Error() string {
	return fmt.Sprintf("%d error %s", u.Code, u.Msg)
}

func NewUserError(code int, msg string) error {
	return UserError{
		Code: code,
		Msg:  msg,
	}
}
