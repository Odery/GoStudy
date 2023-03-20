package main

import (
	"fmt"
)

type celsius float32
type kelvin float32
type sensor = func() kelvin
type getRowF = func(row int) (string, string)

func main() {
	ctok := celsiusToKelvin

	fmt.Println()
	fmt.Println(ctok(22))
	fmt.Printf("%T",ctok)
}

func celsiusToKelvin(c celsius) kelvin{
	return kelvin(c + 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func takesFunc(s sensor, i int) int{

	func(i int){
		fmt.Println(i)
	}(1)

	return 0
}

func drawTable(rows int, r getRowF){

}