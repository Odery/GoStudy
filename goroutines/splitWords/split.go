package main

import (
	"fmt"
	"strings"
)

func main(){
	ch := make(chan string)
	s := "Hello beautiful world! How are you?"

	go stringSpliter(ch, s)
	printSplittedWord(ch)
}

func stringSpliter(up chan string, s string){
	for _, word := range(strings.Fields(s)){
		up <- word
	}
	close(up)
}

func printSplittedWord(down chan string){
	for item := range down{
		fmt.Println(item)
	}
}