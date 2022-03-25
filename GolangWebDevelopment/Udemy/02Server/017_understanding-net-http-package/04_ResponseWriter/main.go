package main

import (
	"fmt"
	"net/http"
	"time"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Erdem-Key", "this is from ERDEM")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", "40")
	w.Header().Set("Last-Modified", time.Now().UTC().String())
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
