package main

import (
	"image"
	"log"
	"time"
)

type command int

const(
	right = command(0)
	left = command(1)
)

type RoverDriver struct{
	commandc chan command
}
func main() {
	r := NewRoverDriver()
	time.Sleep(2 * time.Second)
	r.Left()
	time.Sleep(1 * time.Second)
	r.Right()
	time.Sleep(5 * time.Second)
}

func NewRoverDriver() *RoverDriver{
	r := &RoverDriver{
		commandc: make(chan command),
	}

	go r.drive()
	return r
}

func (r *RoverDriver) drive(){
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}

	updateInterval := 250 * time.Millisecond
	next := time.After(updateInterval)

	for{
		select {
		case c:= <- r.commandc:
			switch c{
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("New direction: %v", direction)

			//Change the code so that the delay time gets a half a 
			//second longer with each move
		case <- next:
			pos = pos.Add(direction)
			log.Printf("Moved to the: %v", pos)
			updateInterval += 500 * time.Millisecond
			next = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) Left(){
	r.commandc <- left
}

func (r *RoverDriver)Right(){
	r.commandc <- right
}