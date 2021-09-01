package services

import (
	"missrv/utils"
	"strconv"
)

type UserServiceInf interface {
	GetUserName(uid int64) string
	GetAddr() string
}

type UserService struct{}

func (u UserService) GetUserName(uid int64) string {
	return "OJ8K"
}

func (u UserService) GetAddr() string {
	return strconv.Itoa(utils.ServicePort)
}
