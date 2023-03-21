package main

import (
	"fmt"
)

func main() {
	var arr [5]string

	arr = [5]string{"Neptun","Zeus","Hades"}

	planets := [...]string{ 
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune", 
	   }
	fmt.Printf("%#v\n", arr)
	fmt.Println(planets)
}