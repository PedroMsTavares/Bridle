package main

import (
	"net/http"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := http.NewServeMux()

	server := &http.Server{
		Addr:         ":8443",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	router.HandleFunc("/", Validate)
	router.HandleFunc("/healthz", health)
	go func() {
		err := server.ListenAndServeTLS("/certificates/cert.pem", "/certificates/key.pem")
		CheckIfError(err)
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("☠️ Shutting down")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	server.Shutdown(c)

}
