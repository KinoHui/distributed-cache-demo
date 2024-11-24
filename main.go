package main

import (
	"distributed-cache-demo/jincache"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *jincache.Group {
	return jincache.NewGroup("scores", 2<<10, jincache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, jin *jincache.Group) {
	peers := jincache.NewHTTPPool(addr)
	peers.Set(addrs...)
	jin.RegisterPeers(peers)
	log.Println("jincache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, jin *jincache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := jin.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil)) // http.ListenAndServe，这个函数会在调用后进入一个阻塞状态，持续监听传入的 HTTP 请求。它不会返回，直到出现错误（例如，端口被占用）或者程序被强制终止。

}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "jincache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	jin := createGroup()
	if api {
		go startAPIServer(apiAddr, jin)
	}
	startCacheServer(addrMap[port], []string(addrs), jin)
}
