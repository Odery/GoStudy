package main

import (
	"fmt"
)

func reset(a *[10]int){
	for i := range a{
		a[i] = 0
	}
}

func main() {
	fmt.Println()

	arr := [10]int{0,1,2,3,4,5,6,7,8,9}

	reset(&arr)

	fmt.Println(arr)
}