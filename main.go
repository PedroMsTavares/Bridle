package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", Validate)
	http.HandleFunc("/ping", ping)
	err := http.ListenAndServeTLS(":8443", "/certificates/cert.pem", "/certificates/key.pem", nil)
	CheckIfError(err)

}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Pong."))
}
