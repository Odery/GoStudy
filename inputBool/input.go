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

	for _, e := range(inputs){
		switch e{
		case "true":
			fmt.Printf("The %v is True\n", e)
		case "yes":
			fmt.Printf("The %v is True\n", e)
		case "1":
			fmt.Printf("The %v is True\n", e)
		case "false":
			fmt.Printf("The %v is False\n", e)
		case "no":
			fmt.Printf("The %v is False\n", e)
		case "0":
			fmt.Printf("The %v is False\n", e)
		default:
			fmt.Printf("The %v is not a bool!\n",e)

		}
	}
}