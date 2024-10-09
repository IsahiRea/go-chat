package main

import (
	"log"
	"net/http"
)

func main() {

	// HTTP handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handlerWebSocket)
	mux.HandleFunc("/healthz", handlerReadiness)

	server := &http.Server{
		Addr:    "8080",
		Handler: mux,
	}

	log.Println("Server started at :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start error: %v", err)
	}
}
