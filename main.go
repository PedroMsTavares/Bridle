package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Validate).Methods("POST")

	http.ListenAndServe(":80", r)
}
