package main

import (
	"image"
	"log"
	"time"
	"sync"
)

type command int

const(
	right = command(0)
	left = command(1)
)

type RoverDriver struct{
	mu sync.Mutex
	commandc chan command
	move bool
}
func main() {
	r := NewRoverDriver()
	r.Left()
	time.Sleep(2 * time.Second)
	
	r.Start()
	time.Sleep(3 * time.Second)

	r.Right()
	time.Sleep(2 * time.Second)

	r.Stop()
	time.Sleep(time.Second)
}

func NewRoverDriver() *RoverDriver{
	r := &RoverDriver{
		commandc: make(chan command),
		move: false,
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

		case <- next:
			if r.move{
				pos = pos.Add(direction)
				log.Printf("Moved to the: %v", pos)
			}else{
				log.Println("Rover is stopped")
			}
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

func (r *RoverDriver)Start(){
	r.mu.Lock()
	defer r.mu.Unlock()
	r.move = true
}

func (r *RoverDriver)Stop(){
	r.mu.Lock()
	defer r.mu.Unlock()
	r.move = false
}