package main

import (
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"net/http"
)

var r = rate.NewLimiter(1, 5)
func MyLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !r.Allow() {
			http.Error(writer, "too many request", http.StatusTooManyRequests)
		} else {
			next.ServeHTTP(writer, request)
		}
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		writer.Write([]byte(`{"status":"OK"}`))
	})
	http.ListenAndServe(":8989", MyLimit(router))
}
