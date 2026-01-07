package main

import (
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Write([]byte("POST request"))
	case http.MethodPut:
		w.Write([]byte("PUT request"))
	case http.MethodDelete:
		w.Write([]byte("DELETE request"))
	case http.MethodGet:
		fallthrough
	default:
		w.Write([]byte("GET request"))
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
}
