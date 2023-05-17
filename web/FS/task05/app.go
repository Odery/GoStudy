package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./web/FS/task05/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("./web/FS/task05/public"))

	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
