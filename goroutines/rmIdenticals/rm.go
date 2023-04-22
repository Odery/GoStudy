package main

import (
	"fmt"
)

func main() {
	down := make(chan string)
	up := make(chan string)
	input := []string{"Bobby", "Jhon", "Bobby", "Jack", "Thomas", "Jack"}

	go sourceGopher(down, input)
	go filterGopher(down, up)
	printGopher(up)
}

func sourceGopher(upStream chan string, input []string){
	for _, v := range input{
		upStream <- v
	}
	close(upStream)
}

//Its to boring to check if the last sent item is equal to current
//Much more practical is to check whether the item was sent previously
//So, it does more than asked
func filterGopher(donwStream, upStream chan string){
	sentElements := []string{""}
	for item := range donwStream{
		canBeSend := true
		for _, v := range sentElements{
			if item == v{
				canBeSend = false
			}
		}
		if canBeSend{
			upStream <- item
			sentElements = append(sentElements, item)
		}
	}
	close(upStream)
}

func printGopher(upStream chan string){
	for item := range upStream{
		fmt.Println(item)
	}
}