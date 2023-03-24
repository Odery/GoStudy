package main

import (
	"fmt"
)

const (
	width = 80
	height = 10
)

type universe [][]bool

func (u universe) show(){
	for _,row := range(u){
		for _, cell := range(row){
			if cell{
				fmt.Print("*")
			}else{
				fmt.Print("/")
			}
		}
		fmt.Println()
	}
}

func main(){
	uni := newUnivese()
	uni.show()
}

func newUnivese() universe{
	uni := make(universe, height)
	for i := range(uni){
		uni[i] = make([]bool, width)
	}
	return uni
}
