package main

import (
	. "github.com/qingsong-he/ce"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func Case1() {

	go func() {
		bk1, _ := url.Parse("http://localhost:30001")
		bk2, _ := url.Parse("http://localhost:30002")
		rp1 := httputil.NewSingleHostReverseProxy(bk1)
		rp2 := httputil.NewSingleHostReverseProxy(bk2)

		rand.Seed(time.Now().Unix())
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if rand.Int()%2 == 0 {
				rp1.ServeHTTP(w, r)
			} else {
				rp2.ServeHTTP(w, r)
			}
		})

		err := http.ListenAndServe(":3000", nil)
		CheckError(err)
	}()

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("/hello by srv1"))
		})
		err := http.ListenAndServe(":30001", mux)
		CheckError(err)
	}()

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("/hello by srv2"))
		})
		err := http.ListenAndServe(":30002", mux)
		CheckError(err)
	}()

	select {}
}

func main() {
	Case1()
	Print("main end")
}
