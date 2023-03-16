/*
Write a program with celsius, fahrenheit, and kelvin types and methods to convert from 
any temperature type to any other temperature type. 
*/

package main

import (
	"fmt"
)

type celsius float32
type kelvin float32
type fahrenheit float32

func main(){
	var tempC celsius = 32
	var tempK kelvin = 320
	var tempF fahrenheit = 120
	fmt.Printf("%.5v C = %.5v K = %.5v F\n",tempC,tempC.kelvin(),tempC.fahrenheit())
	fmt.Printf("%.5v K = %.5v C = %.5v F\n",tempK,tempK.celsius(),tempK.fahrenheit())
	fmt.Printf("%.5v F = %.5v K = %.5v C\n",tempF,tempF.kelvin(),tempF.celsius())
}

func (k kelvin) celsius() celsius{
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit{
	return fahrenheit((k - 273.15) * 9.0/5.0 + 32.0)
}

func (c celsius) fahrenheit() fahrenheit{
	return fahrenheit(c * 9.0 / 5.0 + 32.0)
}

func (c celsius) kelvin() kelvin{
	return kelvin(c + 273.15)
}

func (f fahrenheit) celsius() celsius{
	return celsius((f - 32) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin{
	return kelvin((f - 32) * 5.0 / 9.0 + 273.15)
}