package main

import (
	"fmt"
	"math/rand"
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
				fmt.Print(" ")
			}12
		}
		fmt.Println()
	}
}

func (u universe) seed(){
	for _, row := range(u){
		for i := range(row){
			if rand.Intn(100) <= 25{
				row[i] = true
			}
		}
	}
}

func main(){
	uni := newUnivese()
	uni.seed()
	uni.show()
}

func newUnivese() universe{
	uni := make(universe, height)
	for i := range(uni){
		uni[i] = make([]bool, width)
	}
	return uni
}
