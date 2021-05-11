package biu

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	defaultCachePath = "_cache"
)

type CacheHttpHandler struct {
}

func NewCacheHttpHandler() *CacheHttpHandler {
	return &CacheHttpHandler{}
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("Internal Server Error"))
}
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("NOT FOUND"))
}

func (h *CacheHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	//第一个'/'之前的为空
	if len(split) == 0 || len(split) != 4 || split[1] != defaultCachePath {
		http.Error(w, fmt.Sprint("params error"), http.StatusNotFound)
		return
	}
	groupName, key := split[2], split[3]
	g := getGroup(groupName)
	if g == nil {
		http.Error(w, fmt.Sprintf("group not found with [%s]", groupName), http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodGet {
		if value, err := g.Get(key); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write(value.ByteSlice())
		}
	}

}

func StartServe() {
	if err := http.ListenAndServe(":8082", NewCacheHttpHandler()); err != nil {
		panic(err)
	}
}
