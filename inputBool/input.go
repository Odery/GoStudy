/*
Write a program that converts strings to Booleans:
 The strings “true”, “yes”, or “1” are true.
 The strings “false”, “no”, or “0” are false.
 Display an error message for any other values.
*/

package main

import(
	"fmt"
)

func main(){
	inputs := [8]string{"true","yes","1","false","no","0","gibrisch","Hello World"}
	var launch [8]bool

	for i, e := range(inputs){
		switch e{
		case "true", "yes", "1":
			fmt.Printf("The %v is True\n", e)
			launch[i] = true
		case "false", "no", "0":
			fmt.Printf("The %v is True\n", e)
			launch[i] = false
		default:
			fmt.Printf("The %v is not a bool!\n",e)
			launch[i] = false
		}
	}
	fmt.Println(launch)
}