package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func articleList(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "articleService: articleList")
}
func articleRetrieve(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "articleService: articleRetrieve")
}
func articleCreate(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "articleService: articleCreate")
}
func articleDelete(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "articleService: articleDelete")
}
func articleUpdate(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "articleService: articleUpdate")
}
func articleUpdatePartial(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "articleService: articleUpdatePartial")
}

func main() {
	router := mux.NewRouter()

	// rest api 定义
	router.HandleFunc("/articles", articleList).Methods("GET")
	router.HandleFunc("/articles/{id}", articleRetrieve).Methods("GET")
	router.HandleFunc("/articles", articleCreate).Methods("POST")
	router.HandleFunc("/articles/{id}", articleDelete).Methods("DELETE")
	router.HandleFunc("/articles/{id}", articleUpdate).Methods("PUT")
	router.HandleFunc("/articles/{id}", articleUpdatePartial).Methods("PATCH")

	// http监听
	log.Fatalln(http.ListenAndServe(":8088", router))
}
