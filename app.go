package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Error types declaration
var(
	ErrBound = errors.New("Out of bound")
	ErrDigit = errors.New("Invalid digit")
)
const rows, columns = 9, 9

type SudokuError []error

func (se SudokuError) Error() string{
	var s []string
	for _, err := range se{
		s = append(s, err.Error())
	}

	return strings.Join(s, ", ")
}
type grid [rows][columns]int8

func (g *grid) Set(row,column int8, digit  int8) error{
	var errors SudokuError
	if !inBounds(row, column) {
		errors = append(errors, ErrBound)
	}
	if !validDigit(digit){
		errors = append(errors, ErrDigit)
	}
	if len(errors) > 0{
		return errors
	}

	g[row][column] = digit
	return nil
}

func main() {
	var g grid
	err := g.Set(6,90,85)
	if err != nil{
		if errs, ok := err.(SudokuError); ok{
			fmt.Printf("%d error(s) occurred: \n", len(errs))
			for _, err := range errs{
				fmt.Printf("- %v\n", err)
			}
		}
		os.Exit(1)
	}
	fmt.Println()

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