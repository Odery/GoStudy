package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><head><title>Index</title></head><body><h1>Index Page</h1></body></html>")
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><head><title>Dog</title></head><body><h1>Dog Page</h1></body></html>")
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><head><title>Me</title></head><body><h1>I am Steve!</h1></body></html>")
}