package main

import (
	"fmt"
)

func main() {
	fmt.Println("")

	m := &map[string]int{}

	s := &[3]string{}

	(*m)["bob"] = 3

	s[1] = "Bob"


	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", s)
}