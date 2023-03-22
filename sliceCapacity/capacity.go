package main

import(
	"fmt"
)

func main(){
	slice := make([]int, 0)
	c := cap(slice)
	fmt.Printf("Initial cap size is: %v\n", c)
	for i := 0; i < 500; i++{
		slice = append(slice, i)
		if c < cap(slice){
			c = cap(slice)
			fmt.Printf("Slice changed, new cap is: %v\n", c)
		}
	}
}