package main

import (
	"errors"
	"fmt"
	"strings"
	"math/rand"
)

// Error types declaration
var(
	ErrBound = errors.New("Out of bound")
	ErrDigit = errors.New("Invalid digit")
	ErrNotEmpty = errors.New("The cell contains a value already")
	ErrNotValid = errors.New("Value provided is not valid Sudoku value for that cell")
	ErrCantBeOverwriten = errors.New("The cell can not be changed")
)
const rows, columns = 9, 9

type SudokuError []error
type SudokuCell struct{
	value int8
	notErasable bool
}

func (s SudokuCell) String() string{
	return fmt.Sprint(s.value)
}
type SudokuBracket struct{
	startRow, endRow int8
	startCol, endCol int8
}
func (se SudokuError) Error() string{
	var s []string
	for _, err := range se{
		s = append(s, err.Error())
	}

	return strings.Join(s, ", ")
}
type grid [rows][columns]SudokuCell

func NewGrid() grid{
	var g grid
	for r := range g{
		for c := range g[r]{
			if rand.Intn(100) < 35{
				g.Set(int8(r),int8(c),int8(rand.Intn(10) + 1), true)
			}else{
				g[r][c] = SudokuCell{0, false}
			}
		}
	}
	return g
}

func (g grid) Show() {
    divider := strings.Repeat("=", 31)
    fmt.Println(divider)
    
    for _, row := range g {
        fmt.Print("|")
        for _, cell := range row {
            fmt.Printf("%-3v", cell)
        }
        fmt.Println("|")
    }
    
    fmt.Println(divider)
}

func (g *grid) Erase(row,column,digit int8) bool{
	if g[row][column].notErasable{
		return false
	}else {
		g[row][column].value = 0
		return true 
	}
}

func (g *grid) Set(row,column int8, digit  int8, permanent bool) (bool,error){
	var errors SudokuError
	if !inBounds(row, column) {
		errors = append(errors, ErrBound)
	}
	if !validDigit(digit){
		errors = append(errors, ErrDigit)
	}
	if g[row][column].value != 0{
		errors = append(errors, ErrNotEmpty)
	}
	if g[row][column].notErasable == true{
		errors = append(errors, ErrCantBeOverwriten)
	}
	if g.isValidAnswer(row,column,digit){
		g[row][column] = SudokuCell{digit, permanent}
	}else{
		errors = append(errors, ErrNotValid)
	}
	if len(errors) > 0{
		return false, errors
	}

	return true, nil
}

//It wont solve it completele, needs bactracking algorithm to do so
func (g *grid) Solve(){
	for r := range g{
		for c := range g[r]{
			if g[r][c].value == 0{
				for i:= 1; i <= 10; i++{
					if ok,_ := g.Set(int8(r),int8(c),int8(i),false); ok{
						break
					}
				}
			}
		}
	}
}

func main() {
	g := NewGrid()
	g.Show()
	g.Solve()
	g.Show()

}

func validDigit(digit int8) bool{
	if digit > 0 && digit < 10{
		return true
	}
	return false
}

func inBounds(row, column int8) bool{
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func (g grid) isValidAnswer(row, column, digit int8) bool{
	//Check whether there is the same digit in column
	for _,e := range g[row]{
		if e.value == digit{
			return false
		}
	}
	//Check whether there is the same digit in row
	for r := range g{
		if g[r][column].value == digit{
			return false
		}
	}
	//Check whether there is the same digit in Sudoku bracket
	//Determine what bracket it is 
	var bLimits SudokuBracket
	switch{
	case column < 3 && row < 3:
		bLimits = SudokuBracket{0,3,0,3}
	case column < 6 && row < 3:
		bLimits = SudokuBracket{0,3,3,6}
	case column < 9 && row < 3:
		bLimits = SudokuBracket{0,3,6,9}
	case column < 3 && row < 6:
		bLimits = SudokuBracket{3,6,0,3}
	case column < 6 && row < 6:
		bLimits = SudokuBracket{3,6,3,6}
	case column < 9 && row < 6:
		bLimits = SudokuBracket{3,6,6,9}
	case column < 3 && row < 9:
		bLimits = SudokuBracket{6,9,0,3}
	case column < 6 && row < 9:
		bLimits = SudokuBracket{6,9,3,6}
	case column < 9 && row < 9:
		bLimits = SudokuBracket{6,9,6,9}
	}

	//To slice a submatrix, we need to iterate through the rows and get the desired columns
	for i:= bLimits.startRow; i < bLimits.endRow; i++{
		for _, e := range g[i][bLimits.startCol : bLimits.endCol]{
			if e.value == digit{
				return false
			}
		}
	}

	return true
}