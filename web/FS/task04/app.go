package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./web/FS/task04/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("./web/FS/task04/public"))

	http.Handle("/resources/", http.StripPrefix("/resources/", fs))
	http.HandleFunc("/", index)

	http.ListenAndServe("localhost:8080", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
