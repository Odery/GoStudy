package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width = 80
	height = 10
)

type universe [][]bool

func (u universe) Show(){
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

func (u universe) Seed(){
	for _, row := range(u){
		for i := range(row){
			if rand.Intn(100) <= 25{
				row[i] = true
			}
		}
	}
}

func (u universe) CellAlive(row, cell int) bool{
	row, cell = _normalize(row, cell)
	return u[row][cell]
}

func (u universe) Neighbors(row,cell int) int{
	row, cell = _normalize(row, cell)
	var aliveMembers int
	for r:= -1; r < 2; r++{
		for i:= -1; i < 2; i++{
			if u.CellAlive(row + r, cell + i){
				aliveMembers++
			}
		}
	}
	//Remove 1 if cell is currently alive (must not count itself)
	if u.CellAlive(row,cell){
		aliveMembers --
	}
	return aliveMembers
}

func (u universe) Next(row,cell int) bool{
	n := u.Neighbors(row,cell)
	if n == 2{
		if u.CellAlive(row,cell){
			return true
		}else{
			return false
		}
	}else if n == 3{
		return true
	}else{
		return false
	}
}

func main(){
	uni := NewUnivese()
	uni.Seed()

	for i := 0; i < 100; i++{
		uni.Show()
		time.Sleep(time.Second)
		uni = NextGeneration(uni)
		fmt.Println("\033[2J")
	}
}

func NewUnivese() universe{
	uni := make(universe, height)
	for i := range(uni){
		uni[i] = make([]bool, width)
	}
	return uni
}

func NextGeneration(u universe) universe{
	newUni := NewUnivese()
	for r := range(u){
		for c := range(u[r]){
			newUni[r][c] = u.Next(r,c)
		}
		fmt.Println()
		
	}

	return newUni
}

func _normalize(row, cell int) (int,int){
	return ((row + height) % height), ((cell + width) % width)
}
