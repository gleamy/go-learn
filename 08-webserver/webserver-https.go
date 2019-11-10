package main

// - 安全文件及私钥问题本示例未完结
import (
	"fmt"
	"net/http"
	"log"
//	"time"
	"html"
)

func main() {
//	svr := &http.Server {
//		Addr: ":10443",
//		ReadTimeout: 10 * time.Second,
//		WriteTimeout: 10 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//		Handler: func(w http.ResponseWriter, r *http.Request) {
//			fmt.Fprintf(w, "Hello, &q", html.EscapeString(r.URL.Path))
//		},
//	}
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, &q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServeTLS(":10443", "my.pem", "my.key", nil))
}
