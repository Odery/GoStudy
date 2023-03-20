package main

import (
	"fmt"
)

type celsius float32
func (c celsius) toFahrenheit() fahrenheit{
	return fahrenheit(c * 9.0 / 5.0 + 32.0)
}

type fahrenheit float32
func (f fahrenheit) toCelsius() celsius{
	return celsius((f - 32) * 5.0 / 9.0)
}

type getRow func(row int) (float32, float32)

const (
	line = "======================="
	rowFormat = "| %-8v | %-8v |\n"
	numberFormat = "%.1f"
)

func main(){
	drawTable(getCelsiusRow)
	drawTable(getFahrenheitRow)
}

func drawTable(row getRow){
	fmt.Println(line)
	v,_ := row(0)  
	if v == 0{
		fmt.Printf(rowFormat, "ºC", "ºF")
	}else{
		fmt.Printf(rowFormat, "ºF", "ºC")
	}
	fmt.Println(line)
	fmt.Printf(rowFormat, "-40", "-40")

	for i := 1; i <= 28; i++{
		a,b := row(i)
		fmt.Printf(rowFormat,fmt.Sprintf(numberFormat,a),fmt.Sprintf(numberFormat,b))
	} 
	fmt.Println(line)
}

func getCelsiusRow(row int) (float32, float32){
	var c celsius
	if row != 0 {
		c = celsius(-40 + row * 5)
	}

	return float32(c), float32(c.toFahrenheit())
}

func getFahrenheitRow(row int) (float32, float32){
	var f fahrenheit = 32
	if row != 0 {
		f = fahrenheit(-40 + row * 5)
	}

	return float32(f), float32(f.toCelsius())
}