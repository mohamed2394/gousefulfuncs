package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	errorResponse := map[string]string{"error": msg}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorResponse)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
