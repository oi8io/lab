package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/time/rate"
	"missrv/services"
	"missrv/utils"
)

type UserRequest struct {
	Uid int64 `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
	Addr   string `json:"addr"`
}



func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, utils.NewUserError(429, " Too many requests")
			}
			return next(ctx, request)
		}
	}
}



func UserServiceLoggerMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r:=request.(UserRequest)
			logger.Log("Method","GET","UserID",r.Uid)
			return next(ctx, request)
		}
	}
}

func GenUserEndpoint(service services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)
		username := service.GetUserName(req.Uid)
		return UserResponse{Result: username, Addr: service.GetAddr()}, nil
	}
}
