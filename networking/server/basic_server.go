package server

import (
	"net/http"
)

func Server(){
	http.Handle("/", &home{})
	http.Handle("/blog", &blog{})
	http.ListenAndServe(":8080", nil)
}

type home struct{}
type blog struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("This is my home"))
}

func (b *blog) ServeHTTP(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("This is my blog"))
}

