package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is up and running!"))
}
