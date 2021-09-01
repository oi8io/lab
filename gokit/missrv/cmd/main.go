package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	klog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"log"
	"missrv/endpoints"
	"missrv/services"
	"missrv/transports"
	"missrv/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func checker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"status":"OK"}`))
}

var limit = rate.NewLimiter(1, 5)
var logger klog.Logger

func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			contentType, body = "application/json; charset=utf-8", jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	code := http.StatusTooManyRequests
	if ue, ok := err.(utils.UserError); ok {
		w.WriteHeader(ue.Code)
		w.Write([]byte(ue.Msg))
	} else {
		w.WriteHeader(code)
		w.Write(body)
	}
}

func getHandler() *mux.Router {
	user := services.UserService{}
	{
		logger = klog.NewLogfmtLogger(os.Stdout)
		logger = klog.WithPrefix(logger, "time", klog.DefaultTimestampUTC)
	}
	endpoint := endpoints.UserServiceLoggerMiddleware(logger)(endpoints.RateLimit(limit)(endpoints.GenUserEndpoint(user)))
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(MyErrorEncoder),
	}
	handle := kithttp.NewServer(endpoint, transports.DecodeUserRequest, transports.EncodeUserResponse, options...)
	r := mux.NewRouter()
	//r.Handle("/user/{uid:\\d+}", handle)
	r.Methods("GET", "DELETE").Path("/user/{uid:\\d+}").Handler(handle)
	r.Methods("GET").Path("/health").HandlerFunc(checker)
	//r.Methods("GET", "POST", "HEAD", "PUT", "DELETE", "OPTIONS", "TRACE")
	return r
}

func main() {
	var name = flag.String("name", "", "服务名称")
	var port = flag.Int("port", 0, "端口号")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*port)
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	r := getHandler()
	errChan := make(chan error)
	go func() {
		if err := utils.RegistryService(*name, *port); err != nil {
			log.Fatal("RegistryService", err)
		}
		if err := http.ListenAndServe(addr, r); err != nil {
			errChan <- err
			log.Fatal("ListenAndServe error", err)
		}
	}()
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM) // 监听中断interrupted和终止信号terminated
		errChan <- fmt.Errorf("%s", <-sig)
	}()
	getErr := <-errChan
	if getErr != nil {
		_ = utils.DeregisterService()
	}
}
