package main

import "fmt"

type coordinate struct{
	degrees,minutes,seconds float64
	hemisphere rune
}

func (c coordinate) decimal() float64{
	sign := 1.0
	switch c.hemisphere{
	case 'S','W','s','w':
		sign = -1.0
	}
	return sign * (c.degrees + c.minutes/60 + c.seconds/3600)
}

type location struct{
	lat,lon float64
}
func main(){
	spiritLat := coordinate{degrees: 14,minutes: 34,seconds: 6.2,hemisphere: 'S'}
	spiritLon := coordinate{degrees: 175,minutes: 28,seconds: 21.5,hemisphere: 'E'}

	oppoLat := coordinate{degrees: 1,minutes: 56,seconds: 46.3,hemisphere: 'S'}
	oppoLon := coordinate{degrees: 354,minutes: 28,seconds: 24.2,hemisphere: 'E'}

	curiosityLat := coordinate{degrees: 4,minutes: 35,seconds: 22.2,hemisphere: 'E'}
	curiosityLon := coordinate{degrees: 137,minutes: 26,seconds: 30.1,hemisphere: 'E'}

	inSightLat := coordinate{degrees: 4,minutes: 30,seconds: 0,hemisphere: 'N'}
	inSightLon := coordinate{degrees: 135,minutes: 54,seconds: 0,hemisphere: 'E'}

	locationMap := make(map[string] location)

	locationMap["spirit"] = NewLocation(spiritLat,spiritLon)
	locationMap["opportunity"] = NewLocation(oppoLat, oppoLon)
	locationMap["curiosity"] = NewLocation(curiosityLat, curiosityLon)
	locationMap["inSight"] = NewLocation(inSightLat, inSightLon)

	for name, loc := range(locationMap){
		fmt.Printf("Name: %s, coordinates in decimal: %.5f\n", name, loc)
	}
}

func NewLocation(lat, lon coordinate) location{
	return location{lat.decimal(), lon.decimal()}
}