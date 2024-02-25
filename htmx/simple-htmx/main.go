package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/clicked", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "<p>You clicked the button!</p>")
	})

	http.HandleFunc("/mouse_entered", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "<p>Here mouse mouse!</p>")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "helloooo")
	})

	http.ListenAndServe(":8080", nil)
}

