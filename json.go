package main

import (
	"bytes"
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

// IntOrString is an int that may be unmarshaled from either a JSON number
// literal, or a JSON string. // literal, or a JSON string.
type IntOrString int

func (i *IntOrString) UnmarshalJSON(d []byte) error {
	var v int
	err := json.Unmarshal(bytes.Trim(d, `"`), &v)
	*i = IntOrString(v)
	return err
}

// SliceOrString is a slice of strings that may be unmarshaled from either // SliceOrString is a slice of strings that may be unmarshaled from either
// a JSON array of strings, or a single JSON string. // a JSON array of strings, or a single JSON string.
type SliceOrString []string

func (s *SliceOrString) UnmarshalJSON(d []byte) error {
	if d[0] == '"' {
		var v string
		err := json.Unmarshal(d, &v)
		*s = SliceOrString{v}
		return err
	}
	var v []string
	err := json.Unmarshal(d, &v)
	*s = SliceOrString(v)
	return err
}

type Success struct {
	Results []string `json:"results"`
}

type Error struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

type Response struct {
	Success
	Error
}
