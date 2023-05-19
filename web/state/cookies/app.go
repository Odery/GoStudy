package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

type HtmlData struct {
	Title  string
	Msg    string
	Cookie string
}

func init() {
	tpl = template.Must(template.ParseFiles("./web/state/cookies/templates/index.gohtml"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/cookie", cookie)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func index(w http.ResponseWriter, r *http.Request) {
	data := HtmlData{"Main page", "Hello there traveler!", ""}

	c, err := r.Cookie("Visits")
	if err != nil {
		c = &http.Cookie{Name: "Visits", Value: "0"}
	}
	oldValue, _ := strconv.Atoi(c.Value)
	c.Value = strconv.Itoa(oldValue + 1)

	http.SetCookie(w, c)
	err = tpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func cookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Visits")
	var cookieValue string
	if err != nil {
		cookieValue = "No visits yet"
	} else {
		cookieValue = c.Name + ": " + c.Value
	}
	data := HtmlData{"Cookie Browser", "Wanna see a cookie?", cookieValue}

	err = tpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
