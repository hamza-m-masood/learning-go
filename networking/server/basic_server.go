package server

import (
	"fmt"
	"net/http"
)

//TODO Add a database
func Server() {
	http.HandleFunc("/", getRequest)
	http.HandleFunc("/post", postRequest)
	http.HandleFunc("/delete", deleteRequest)
	http.ListenAndServe(":8080", nil)
}

func getRequest(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/"{
		errorHandler(w, req, 404)
		return
	}
	if req.Method != http.MethodGet {
		http.Error(w, "Incorrect method", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "You just made a get request")
}

//TODO: Accept json from the client
func postRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Incorrect method", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "You just made a post request")
}

func deleteRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Incorrect method", http.StatusInternalServerError)
		return
	}
	// w.Write([]byte("You just made a delete request"))
	fmt.Fprintln(w, "You just made a delete request")
}

func errorHandler(w http.ResponseWriter, _ *http.Request, status int){
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprintln(w, "ERROR: page not found")
	}
}
