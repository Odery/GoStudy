package main

import (
	"fmt"
	"math"
)
type location struct{
	name string
	lat, lon float64
}

type world struct{
	radius float64
}

type gps struct{
	current location
	destination location
	world
}

type rover struct{
	name string
	gps
}

func main(){
	cur := location{name: "Bradbury lending",lat: -4.5895, lon: 137.4417}
	dst := location{name: "Elysium Planitia",lat: 4.5, lon: 135.9}
	mars := world{radius: 3389.5}
	curiosity := rover{name: "Curiosity", gps: gps{current: cur, destination: dst, world: mars }}

	msg := curiosity.message()

	fmt.Println(msg)
	fmt.Printf("Current location: %s\n", curiosity.current.description())
	fmt.Printf("Destination location: %s\n", curiosity.destination.description())
}

func (l location) description() string{
	return fmt.Sprintf("Name: %s, Latitude: %.4f, Longitude: %.4f", l.name,l.lat,l.lon)
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(l1,l2 location) float64{
	sin, cos := math.Sincos(rad(l1.lat))
	sin2, cos2 := math.Sincos(rad(l2.lat))
	clong := math.Cos(rad(l1.lon - l2.lon))
	return w.radius * math.Acos(sin * sin2 + cos * cos2 * clong)
}
func rad(deg float64) float64{
	return deg * math.Pi / 180
}

func (g gps) message() string{
	d := g.distance(g.current, g.destination)
	return fmt.Sprintf("Distance to the destination is: %0.2f Km", d)
}