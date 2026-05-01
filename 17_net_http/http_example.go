package main

import (
	"encoding/json"
	"net/http"
)

type statusResponse struct {
	Message string `json:"message"`
}

func statusHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(statusResponse{Message: "ok"})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Learn", "golang")
		next.ServeHTTP(w, r)
	})
}
