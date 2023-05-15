package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Response struct{
	Title string
	Body string
}

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("./tpl.html")
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{"Index", "This is an index page"}
	tpl.Execute(w, res)
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{"Dog", "This is a dog page"}
	tpl.Execute(w, res)
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{"Myself", "This is a page about me, Steve."}
	tpl.Execute(w, res)
}