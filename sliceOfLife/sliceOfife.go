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
			}
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

func (u universe) cellAlive(row, cell int) bool{
	row, cell = _normalize(row, cell)
	return u[row][cell]
}

func (u universe) neighbors(row,cell int) int{
	row, cell = _normalize(row, cell)
	var aliveMembers int
	for r:= -1; r < 2; r++{
		for i:= -1; i < 2; i++{
			if u.cellAlive(row + r, cell + i){
				aliveMembers++
			}
		}
	}
	//Remove 1 if cell is currently alive
	if u.cellAlive(row,cell){
		aliveMembers --
	}
	return aliveMembers
}

func main(){
	uni := newUnivese()
	uni.show()
	fmt.Println(uni.neighbors(2,2))
}

func newUnivese() universe{
	uni := make(universe, height)
	for i := range(uni){
		uni[i] = make([]bool, width)
	}
	return uni
}

func _normalize(row, cell int) (int,int){
	return ((row + width) % width), ((cell + height) % height)
}
