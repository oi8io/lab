package main

import (
	"fmt"
	"log"
	"oi.io/apps/biu/biu"
)

func main() {
	var db = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}
	loadCounts := make(map[string]int, len(db))
	biu.NewCacheGroup("c1", 2<<10, biu.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			if _, ok := loadCounts[key]; !ok {
				loadCounts[key] = 0
			}
			loadCounts[key] += 1
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))
	biu.StartServe()
}
