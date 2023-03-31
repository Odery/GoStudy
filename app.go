package main

import (
	"fmt"
)

type coordinate struct{
	d,m,s float64
	h rune
}

type location struct{
	name string
	lat, lon coordinate
}

func (l location) String() string{
	return fmt.Sprintf("%s is at: %0.2fº%0.2f'%0.2f\" %c, %0.2fº%0.2f'%0.2f\" %c", l.name, l.lat.d, l.lat.m, l.lat.s, l.lat.h, l.lon.d,l.lon.m, l.lon.s, l.lon.h)
}
func main() {
	l := location{name: " Elysium Planitia", lat: coordinate{d: 4, m: 30, h: 'N'}, lon: coordinate{d: 135, m: 54, h: 'E'} }
	fmt.Println(l)
}