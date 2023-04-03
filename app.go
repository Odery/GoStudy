package main

import (
	"fmt"
)

type celsius float64

func main() {
	fmt.Println()

	m := make(map[int] interface{})

	m[1] = true
	m[2] = 23.1
	m[3] = 44
	m[4] = celsius(20.9)
	m[5] = 1 + 3i

	for i, v := range m{
		switch e := v.(type){
		case celsius:
			fmt.Printf("Element %v is value of celsius: %v\n", i, e)
		case int:
			fmt.Printf("Element %v is value of int: %v\n", i, e)
		case float64:
			fmt.Printf("Element %v is value of float: %v\n", i, e)
		case string:
			fmt.Printf("Element %v is value of string: %v\n", i, e)
		case bool:
			fmt.Printf("Element %v is value of bool: %v\n", i, e)
		case complex128:
			fmt.Printf("Element %v is value of complex: %v\n", i, e)
		}
	}
}