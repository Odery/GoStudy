package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type cat struct{
	N string `json:"Name"`
	B string `json:"Breed"`
	A uint8` json:"Age"`
}

var car struct{
	Brand string
	Model string
	Age uint8
}

func main() {
	cats := []cat{
		{"Tom","White",2},
		{"George","Britain",9},
		{"Igor", "Local", 3},
	}

	b, err := json.Marshal(cats)
	exitOnErr(err)
	fmt.Println(string(b))
}

func exitOnErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
