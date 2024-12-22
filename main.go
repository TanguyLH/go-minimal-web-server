package main

import (
	"fmt"
	"net/http"
)


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

func mySetupServer() *Router {
	r := &Router{}
	return r
}

func main() {
	r := mySetupServer()
	http.ListenAndServe(":8080", r)
	fmt.Println("> Getting started ... ")
}
