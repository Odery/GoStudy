package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	//Checking whether request is for dog pic for this binding
	//if so, server dog pic, and return
	if r.URL.Path == "/dog/german_shepard.png" {
		http.ServeFile(w, r, "./web/FS/task01/resources/german_shepard.png")
		return
	}
	tpl, err := template.ParseFiles("./web/FS/task01/dog.gohtml")
	if err != nil {
		log.Println(err)
	}

	log.Println(tpl.Execute(w, nil))
}

func foo(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "foo ran")
	if err != nil {
		log.Println(err)
	}
}
