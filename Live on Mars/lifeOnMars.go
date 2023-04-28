package main

import (
	"image"
	"sync"
	"time"
	"math/rand"
)

type command int

const (
	right = command(0)
	left = command(1)
)

func main() {

}

//Rover type to run around the MarsGrid
type Rover struct{
	mu sync.Mutex
	command chan command
	mg *MarsGrid
	move bool
}

//Constructor for a new Rover
func NewRover(mg *MarsGrid) *Rover{
	r := Rover{
		command: make(chan command),
		move: false,
		mg: mg,
	}

	go r.Drive()
	return &r
}
//Drive func for Rover to start movement
func (r *Rover) Drive(){
	next := time.After(250 * time.Millisecond)
	for {
		select{
		case <- next:

		}
	}
}

// MarsGrid represents a grid of some of the surface
// of Mars. It may be used concurrently by different
// goroutines.
// It is pointless to add mutex field to whole Grid,
// so the grid cell will have it instead
type MarsGrid struct{
	grid [100][100]Cell
}

func NewMarsGrid() *MarsGrid{
	m := MarsGrid{}

	for i := range m.grid{
		for j := range m.grid[i]{
			m.grid[i][j] = Cell{
				chanceOfLife: rand.Intn(1000) + 1,
			}
		}
	}

	return &m
}

// Occupy occupies a cell at the given point in the grid. It
// returns nil if the point is already occupied or the point is
// outside the grid. Otherwise it returns a value that can be
// used to move to different places on the grid
func (m *MarsGrid) Occupy(p image.Point) *Occupier{
	//TODO
	return nil
}

// Grid cell to represent state of the grid
type Cell struct{
	mu sync.Mutex
	chanceOfLife int
	Occupier
}

// Occupier represents an occupied cell in the grid.
// It may be used concurrently by different goroutines.
type Occupier struct{
	mu sync.Mutex
	occupied bool
}

// Move moves the occupier to a different cell in the grid.
// It reports whether the move was successful
// It might fail because it was trying to move outside
// the grid or because the cell it's trying to move into
// is occupied. If it fails, the occupier remains in the same place.
func (o *Occupier) Move(p image.Point) bool{
	//TODO
	return false
}
