package biu

import (
	"fmt"
	"net/http"
	"oi.io/apps/biu/biu/consistent"
	"strings"
	"sync"
)

const (
	defaultBasePath = "_cache"
	defaultReplicas = 50
)

// HTTPPool implements PeerPicker for a pool of HTTP peers.
type HTTPPool struct {
	// this peer's base URL, e.g. "https://example.net:8000"
	self        string
	basePath    string
	mu          sync.Mutex // guards peers and httpGetters
	peers       *consistent.HashMap
	httpGetters map[string]*httpGetter // keyed by e.g. "http://10.0.0.2:8008"
}

func NewHTTPPool(self string, ) *HTTPPool {
	return &HTTPPool{self: self, basePath: defaultBasePath, httpGetters: make(map[string]*httpGetter)}
}

func (h *HTTPPool) PickPeer(key string) (peer PeerGetter, ok bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	nodeName, err := h.peers.Get(key)
	if err != nil || nodeName == "" || nodeName == h.self {
		return
	}
	peer, ok = h.httpGetters[nodeName]
	return
}

var _ PeerPicker = (*HTTPPool)(nil)

type CacheHttpHandler struct {
}

// Set updates the pool's list of peers.
func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistent.NewHashMap(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{name: peer, baseURL: peer + "/" + p.basePath + "/"}
	}
}

func (h *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	// 第一个'/'之前的为空
	if len(split) == 0 || len(split) != 4 || split[1] != defaultBasePath {
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
			w.Header().Set("Content-Type", "application/text")
			_, _ = w.Write(value.ByteSlice())
		}
	}
}

func StartServe(h *HTTPPool) {
	if err := http.ListenAndServe(h.self, h); err != nil {
		panic(err)
	}
}
