package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"oi.io/apps/biu/biu"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
	"xxx":  "768",
}

func createGroup(name string) *biu.CacheGroup {
	return biu.NewCacheGroup(name, 2<<10, biu.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

/**
┌─────────────────┐
│     client      │
└─────────────────┘
         │
         ▼
┌─────────────────┐   ┌─────────────────┐                   ┌─────────────────┐
│     server      │   │      peers      │──────────────────▶│      peer       │
└─────────────────┘   └─────────────────┘                   └─────────────────┘
         │                     ▲                                     │
         ▼                     Λ                                     ▼
┌─────────────────┐           ╱ ╲                                    Λ
│      local      │─────────▶▕ S ▏          ┌─────────────────┐     ╱ ╲
└─────────────────┘           ╲ ╱         ┌─│    http get     │◀───▕ s ▏
                               V          │ └─────────────────┘     ╲ ╱
                               │          │                          V
                               ▼          │                          ▼
                          .─────────.     │                 ┌─────────────────┐
                         (    end    )◀───┴─────────────────│       db        │
                          `─────────'                       └─────────────────┘
*/
func startCacheServer(name, addr string, nodeMap map[string]string, gee *biu.CacheGroup) error {
	peers := biu.NewHTTPPool(name)
	peers.Set(nodeMap)
	gee.RegisterPeers(peers)
	log.Println("biu is running at", addr)
	return http.ListenAndServe(addr[7:], peers)
}

func startAPIServer(apiAddr string, c *biu.CacheGroup) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := c.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/text")
			w.Write(view.ByteSlice())
		}))
	log.Println("frontend server is running at", apiAddr, c.GetName())
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))

}

var grr errgroup.Group

func startMultiServers(addrMap map[string]string) {
	apiAddr := "http://localhost:9999"

	var enableApi bool
	for i, v := range addrMap {
		name := i
		addr := v
		gee := createGroup(i)
		if !enableApi {
			enableApi = true
			go startAPIServer(apiAddr, gee)
		}
		grr.Go(func() error {
			return startCacheServer(name, addr, addrMap, gee)
		})
	}
	err := grr.Wait()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Start server success")
}

func main() {
	addrMap := map[string]string{
		"8001": "http://localhost:8001",
		"8002": "http://localhost:8002",
		"8003": "http://localhost:8003",
	}
	startMultiServers(addrMap)
}
