/*
Save some money to buy a gift for your friend. Write a program that randomly places 
nickels ($0.05), dimes ($0.10), and quarters ($0.25) into an empty piggy bank until it
contains at least $20.00. Display the running balance of the piggy bank after each deposit, 
formatting it with an appropriate width and precision. 
*/

package main

import(
	"fmt"
	"math/rand"
)

const nickel = 0.05
const dime = 0.10
const quarter = 0.25
var piggyBank float64

func main(){
	for piggyBank < 20 {
		switch rand.Intn(3) {
		case 0:
			piggyBank += nickel
		case 1:
			piggyBank += dime
		case 2:
			piggyBank += quarter
		}
	} 
	fmt.Println("Piggy-bank contains:")
	fmt.Printf("%12.2f",piggyBank)
}