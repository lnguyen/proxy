package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		remote, err := url.Parse("http://" + r.URL.Path[1:])
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		r.URL.Path = "/"
		proxy.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
