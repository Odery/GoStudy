package main

import (
	"fmt"
)

type celsius float32
type kelvin float32

func main() {
	var temp celsius = 23 
	var temmp kelvin = 300

	fmt.Printf("Type: %T is %[1]v K\n", celsiusToKelvin(temp))
	fmt.Println(temmp.celsius())
}

func celsiusToKelvin(c celsius) kelvin{
	return kelvin(c + 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}