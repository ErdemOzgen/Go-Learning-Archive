package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	http.Handle("/dog", d) //defaultservemux
	http.Handle("/cat", c)
	http.HandleFunc("/", h1)
	http.ListenAndServe(":8080", nil)
}
