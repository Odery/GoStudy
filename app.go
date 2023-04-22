package main

import (
	"fmt"
	"strings"
)

func main() {
	down := make(chan string)
	up := make(chan string)

	go sourceGopher(down)
	go filterGopher(down, up)
	printGopher(up)
}

func sourceGopher(upStream chan string){
	for _, v := range([]string{"hello world", "a bad apple", "goodbye gophers"}){
		upStream <- v
	}
	close(upStream)
}

func filterGopher(donwStream, upStream chan string){
	for item := range donwStream{
		if !strings.Contains(item, "bad"){
			upStream <- item
		}
	}
	close(upStream)
}

func printGopher(upStream chan string){
	for item := range upStream{
		fmt.Println(item)
	}
}