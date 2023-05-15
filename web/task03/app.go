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
	var i index
	var d dog
	var m myself
	http.Handle("/", i)
	http.Handle("/dog/", d)
	http.Handle("/me/", m)

	http.ListenAndServe(":8080", nil)
}

//Handler types for http.Handle
type index int
func (i index) ServeHTTP(w http.ResponseWriter, r *http.Request){
	res := Response{"Index", "This is an index page"}
	tpl.Execute(w, res)
}

type dog int
func (d dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := Response{"Dog", "This is a dog page"}
	tpl.Execute(w, res)
}

type myself int
func (m myself) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := Response{"Myself", "This is a page about me, Steve."}
	tpl.Execute(w, res)
}