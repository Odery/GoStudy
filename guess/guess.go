package main

import (
	"fmt"
	"math/rand"
)
var numToGuess int = 42

func main(){
	trysToGuess := 0
	for{
		guessed := rand.Intn(100) +1
		if guessed == numToGuess{
			fmt.Printf("The number %v was guessed corectly! In %v attempts",numToGuess, trysToGuess)
			break
		}else{
			trysToGuess ++
			if guessed < numToGuess{
				fmt.Println("Guessed number is smaller!")
			}else{
				fmt.Println("Guessed number is bigger!")
			}
		}
	}
}