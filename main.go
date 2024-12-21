package main

import (
	"fmt"
	"net/http"
)

type Router struct {}

func (sr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Return 404 for every request
    http.NotFound(w, r)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Hello World\n"))
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Home\n"))
}

func setupServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getAll) // Putting the handler as a function. It's a delegate ?
	mux.HandleFunc("/home", getHome) 
	return mux
}

func main() {
	server := setupServer()
	http.ListenAndServe(":8080", server)
	fmt.Println("> Getting started ... ")
}
