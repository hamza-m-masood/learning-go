package server

import (
	"fmt"
	"net/http"
)

func Server() {
	//creating a function outside main
	http.HandleFunc("/", homeServe)
	//creating a function inside the params
	http.HandleFunc("/holiday", func(w http.ResponseWriter, _ *http.Request){
		w.Write([]byte("holiday!!"))
		})
	http.ListenAndServe(":8080", nil)
}

func homeServe(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("hello")
	w.Write([]byte("hello"))
}
