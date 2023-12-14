package router

import (
	"fmt"
	"net/http"

)

func NewRouter() http.Handler {
	mux := http.NewServerMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/hello", helloRoute)

	return mux
}

func indexHandler(w http.ResponseHandler, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func helloHander(w http.ResponseHandler, r *http.Request) {
	data := "string of data"
	fmt.Fprintln(w, data)
}