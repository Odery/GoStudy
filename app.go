package main

import (
	"fmt"
)

func main() {

	x := byte(65)
	y := byte(90)

	input_value := byte(63)

	e := (y - x + 1)

	result := ((input_value - x) % e + e) % e + x

	fmt.Printf("%v\t%v\n", result, e)


	deciphered := fmt.Sprintf("%c",((61 - x) % e + e) % e + x )

	fmt.Printf("%v\n", deciphered)

}
