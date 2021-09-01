package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	clog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
	"io"
	"log"
	"miscli/endpoints"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var instance *consul.Instancer
var logger clog.Logger

func init() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	logger = clog.NewLogfmtLogger(os.Stdout)
	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	client := consul.NewClient(consulClient)
	instance = consul.NewInstancer(client, logger, "userService", []string{"primary"}, true)
}

var factory = func(instance string) (endpoint.Endpoint, io.Closer, error) {
	tar, _ := url.Parse("http://" + instance)
	var err error
	return kithttp.NewClient("GET", tar, enc, dec).Endpoint(), nil, err
}

//ConsulEndpoint 通过consul 方式获取
func ConsulEndpoint() (endpoint.Endpoint, error) {
	endpointor := sd.NewEndpointer(instance, factory, logger)
	return lb.NewRandom(endpointor, time.Now().UnixNano()).Endpoint()
	//return lb.NewRoundRobin(endpointor).Endpoint()
}

func GetEndpoint(useConsul bool) (endpoint.Endpoint, error) {
	if useConsul {
		return ConsulEndpoint()
	} else {
		target, err := url.Parse("http://localhost:9999")
		if err != nil {
			return nil, err
		}
		client := kithttp.NewClient("GET", target, enc, dec)
		return client.Endpoint(), nil
	}
}

func DoRequest() {
	endfunc, err := GetEndpoint(true)
	//endfunc, err := ConsulEndpoint()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	res, err := endfunc(ctx, endpoints.UserRequest{Uid: 12})
	if err != nil {
		fmt.Println("get result error", err)
		return
	}
	userInfo := res.(endpoints.UserResponse)
	fmt.Println(time.Now().Format(time.RFC3339), userInfo)
}

func main() {
	for {
		DoRequest()
		time.Sleep(time.Millisecond * 200)
	}
}

func enc(ctx context.Context, req *http.Request, r interface{}) error {
	request := r.(endpoints.UserRequest)
	path := fmt.Sprintf("/user/%d", request.Uid)
	req.URL.Path = path
	return nil
}

func dec(_ context.Context, res *http.Response) (response interface{}, err error) {
	if res.StatusCode > 400 {
		return res, errors.New(strconv.Itoa(res.StatusCode))
	}
	var resp endpoints.UserResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	return resp, nil
}
