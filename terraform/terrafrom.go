package main

import (
	"fmt"
)

type planets []string

func (p planets) terraform() {
	for i, e := range p {
		p[i] = "New " + e
	}
}

func main() {
	var solarPlanetsArray = [8]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	var solarPlanets planets = append(solarPlanetsArray[3:4], solarPlanetsArray[6:8]...)
	solarPlanets.terraform()
	fmt.Println(solarPlanets)
}

