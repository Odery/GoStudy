package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:8080", http.FileServer(http.Dir("./web/FS/task02/starting-files"))))
}
