package main

import (
	"fmt"
	"math/rand"
)

type kelvin float32

type sensor func() kelvin

func realSensor() kelvin{
	return 0
}

func fakeSensor() kelvin{
	return kelvin(rand.Intn(151) + 100)
}

func calibrate(s sensor, offset kelvin) sensor{
	return func() kelvin{
		return s() + offset
	}
}

func main(){
	sensor := calibrate(realSensor, 5)
	fSensor := calibrate(fakeSensor, 500)
	fmt.Println(sensor())
	fmt.Println(fSensor())
}