/* Task:
The table should have four columns:
 The spaceline company providing the service
 The duration in days for the trip to Mars (one-way)
 Whether the price covers a return trip
 The price in millions of dollars

For each ticket, randomly select one of the following spacelines: Space Adventures, 
SpaceX, or Virgin Galactic.
Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km 
away from Earth at the time.
Randomly choose the speed the ship will travel, from 16 to 30 km/s. This will determine 
the duration for the trip to Mars and also the ticket price. Make faster ships more expensive, 
ranging in price from $36 million to $50 million. Double the price for round trips
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const distanceToMars int = 62100000

	var spaceline string
	var duration int
	var returnTicket string
	var price int

	fmt.Printf("%v%12v%12v%8v\n","Spaceline","Days","Trip type","Price")
	fmt.Println("=========================================")

	for i := 0; i < 10; i++ {
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}
		if rand.Intn(100) < 50 {
			returnTicket = "One-way"
		} else {
			returnTicket = "Round-trip"
		}

		//Randomly choose the speed the ship will travel, from 16 to 30 km/s
		randSpeed := rand.Intn(15) + 16

		//calculate how many days trip would take with given speed
		duration = (distanceToMars / randSpeed) / 86400

		//Asuming that progression in price is linear, we can use some Math to find out prices in between
		price = (duration - 23) * (36 - 50) / (44 - 23) + 50

		if returnTicket == "Round-trip" {
			price *= 2
			duration *= 2
		}

		fmt.Printf("%-18v%-6v%-12v$%3v\n", spaceline, duration, returnTicket, price)
	}
}
