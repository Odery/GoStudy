package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./web/FS/task03/templates/index.gohtml"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	fs := http.FileServer(http.Dir("./web/FS/task03/public"))
	mux.Handle("/pics/", fs)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
