package main

import (
	"fmt"
	"math"
)

type coordinate struct{
	degrees,minutes,seconds float64
	hemisphere rune
}

type location struct{
	lat,lon float64
}

type world struct{
	radius float64
}

//Coordinates could be written in a more concise way,
//however, it will compromise readability a lot
var (
	//Landings sites on mars
	spiritLat = coordinate{degrees: 14,minutes: 34,seconds: 6.2,hemisphere: 'S'}
	spiritLon = coordinate{degrees: 175,minutes: 28,seconds: 21.5,hemisphere: 'E'}

	oppoLat = coordinate{degrees: 1,minutes: 56,seconds: 46.3,hemisphere: 'S'}
	oppoLon = coordinate{degrees: 354,minutes: 28,seconds: 24.2,hemisphere: 'E'}

	curiosityLat = coordinate{degrees: 4,minutes: 35,seconds: 22.2,hemisphere: 'E'}
	curiosityLon = coordinate{degrees: 137,minutes: 26,seconds: 30.1,hemisphere: 'E'}

	inSightLat = coordinate{degrees: 4,minutes: 30,seconds: 0,hemisphere: 'N'}
	inSightLon = coordinate{degrees: 135,minutes: 54,seconds: 0,hemisphere: 'E'}

	//Mountains on Mars
	olimpLat = coordinate{degrees: 18, minutes: 39, hemisphere: 'N'}
	olimpLon = coordinate{degrees: 226, minutes: 12, hemisphere: 'E'}

	sharpLat = coordinate{degrees: 5, minutes: 4, seconds: 48, hemisphere: 'S'}
	sharpLon = coordinate{degrees: 137, minutes: 51, hemisphere: 'E'}

	//Cities on Earth
	londonLat = coordinate{degrees: 51, minutes: 30, hemisphere: 'N'}
	londonLon = coordinate{degrees: 0, minutes: 8, hemisphere: 'W'}

	parisLat = coordinate{degrees: 48, minutes: 51, hemisphere: 'N'}
	parisLon = coordinate{degrees: 2, minutes: 21, hemisphere: 'E'}

	kievLat = coordinate{degrees: 50, minutes: 27, seconds: 0.0036, hemisphere: 'N'}
	kievLon = coordinate{degrees: 30, minutes: 31, seconds: 23.9988, hemisphere: 'N'}

	khmelLat = coordinate{degrees: 49, minutes: 25, seconds: 22.74, hemisphere: 'N'}
	khmelLon = coordinate{degrees: 26, minutes: 59, seconds: 13.68, hemisphere: 'N'}

	//Worlds
	mars = world{radius: 3389.5}
	earth = world{radius: 6371.0}
)
//Start of the program
func main(){

	landingsMap := make(map[string] location)
	landingsMap["spirit"] = NewLocation(spiritLat,spiritLon)
	landingsMap["opportunity"] = NewLocation(oppoLat, oppoLon)
	landingsMap["curiosity"] = NewLocation(curiosityLat, curiosityLon)
	landingsMap["inSight"] = NewLocation(inSightLat, inSightLon)

	citiesMap := make(map[string] location)
	citiesMap["London"] = NewLocation(londonLat, londonLon)
	citiesMap["Paris"] = NewLocation(parisLat, parisLon)
	citiesMap["Kiev"] = NewLocation(kievLat, kievLon)
	citiesMap["Khmelnytsky"] = NewLocation(khmelLat, khmelLon)

	mountainsOnMarsMap := make(map[string] location)
	mountainsOnMarsMap["Olimp"] = NewLocation(olimpLat, olimpLon)
	mountainsOnMarsMap["Sharp"] = NewLocation(sharpLat, sharpLon)

	printDistnace(mars, landingsMap)
	printDistnace(mars, mountainsOnMarsMap)
	printDistnace(earth, citiesMap)

}

func printDistnace(w world, locs map[string]location){
	for name, loc := range(locs){
		for name2, loc2 := range(locs){
			if name != name2{
				fmt.Printf("Distance between %s and %s, is: %0.2f km.\n", name, name2, w.distance(loc, loc2))
			}

		}
		fmt.Println()
	}
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(l1,l2 location) float64{
	sin, cos := math.Sincos(rad(l1.lat))
	sin2, cos2 := math.Sincos(rad(l2.lat))
	clong := math.Cos(rad(l1.lon - l2.lon))
	return w.radius * math.Acos(sin * sin2 + cos * cos2 * clong)
}

func (c coordinate) decimal() float64{
	sign := 1.0
	switch c.hemisphere{
	case 'S','W','s','w':
		sign = -1.0
	}
	return sign * (c.degrees + c.minutes/60 + c.seconds/3600)
}

//To use math to calculate distance, we first need to
//convert degrees to radians
func rad(deg float64) float64{
	return deg * math.Pi / 180
}

func NewLocation(lat, lon coordinate) location{
	return location{lat.decimal(), lon.decimal()}
}