package main

import (
	"fmt"
	"strconv"
)

func main() {

	var goopher = "yes"
	var gopher = "true"
	var Gopher = "True"

	fmt.Printf("%v%v%v\n",goopher, gopher, Gopher)

	fmt.Println(strconv.ParseBool(Gopher))
}
