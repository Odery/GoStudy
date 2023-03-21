package main

import (
	"fmt"
)

func main() {
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

	iceGiants := planets[2:4]
	fmt.Println(iceGiants)

	panic("Panicking")

	fmt.Println(iceGiants)
}