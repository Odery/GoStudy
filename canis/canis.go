/*
Canis Major Dwarf is the closest known galaxy to Earth 
at 236,000,000,000,000,000 km from our Sun 
(though some dispute that it is a galaxy). 
Use constants to convert this distance to light years
*/
package main

import(
	"fmt"

)


func main(){
	// This is untyped constant, GO will use math/big for its calculation
	const distanceToCanis = 236000000000000000
	const speedOfLight = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365

	distanceInLightYearsToCanis := ((distanceToCanis / speedOfLight) / secondsPerDay) / daysPerYear

	fmt.Printf("Distance to Canis Major Dwarf Galaxy is: %v light years\n", distanceInLightYearsToCanis)
}