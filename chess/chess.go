package main

import (
	"fmt"
)
type board [8][8]rune

func main(){
	drawBoard(initBoard())
}

func initBoard() board{
	var b board
	b[0] = [8]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'} 
	b[1] = [8]rune{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'}
	for i:= 2; i < 6; i++{
		b[i] = [8]rune{'.', '.', '.', '.', '.', '.', '.', '.'}
	}
	b[6] = [8]rune{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'} 
	b[7] = [8]rune{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'} 

	return b
}

func drawBoard(b board){
	fmt.Println("  A B C D E F G H")
	for i,e := range(b){
		fmt.Print(i + 1)
		for _, ee := range(e){
			fmt.Printf(" %c", ee)
		}
		fmt.Println()
	}
}