package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func tmpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func TestGetCache(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9999/api?key=Tom", nil)
	recorder := httptest.NewRecorder()
	tmpHandler(recorder, request)
	all, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(all))
}

func TestCache(t *testing.T) {
	get, err := http.Get("http://localhost:9999/api?key=Tom")
	if err != nil {
		t.Fatal(err)
	}
	all, _ := ioutil.ReadAll(get.Body)
	fmt.Println(string(all))
}
func BenchmarkCache(b *testing.B) {
	//b.N  = 10
	for i := 0; i < b.N; i++ {
		_, err := http.Get("http://localhost:9999/api?key=Jack")
		if err != nil {
			b.Fatal(err)
		}
	}
}