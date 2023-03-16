/*
 Reuse the kelvinToCelsius function to convert 233 K to Celsius.
 Write and use a celsiusToFahrenheit temperature conversion function.
  Hint: the formula for converting to Fahrenheit is: (c * 9.0 / 5.0) + 32.0.
 Write a kelvinToFahrenheit function and verify that it converts 0 K to approximately –459.67° F.
*/

package main

import (
	"fmt"
)

func main(){
	fmt.Printf("233 K in Celsius is: %.5v\n", kelveinTOCelsius(233))
	fmt.Printf("36 Celsius in Fahrenheit is: %.5v\n", celsiusToFahrenheit(36))
	fmt.Printf("0 K in Farenheit is: %.5v\n", kelvinToFahrenheit(0))

}

func kelveinTOCelsius(k float64) float64{
	return k - 273.15
}

func celsiusToFahrenheit(c float64) float64{
	return c * 9.0 / 5.0 + 32.0
}

func kelvinToFahrenheit(k float64) float64{
	return k - 273.15 * 9.0/5.0 + 32.0
}