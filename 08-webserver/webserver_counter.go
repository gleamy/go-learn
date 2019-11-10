package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
//	hd := func(w http.ResponseWriter, r *http.Request) {
//		lissajous(w)
//	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
//	http.HandleFunc("/lissajous", hd)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// - 回显请求的URL路径部分i
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// - 回显调用次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
