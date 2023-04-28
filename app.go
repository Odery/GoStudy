package main

import (
	"log"
	"sync"
)

func main(){
	log.Println("time.Millisecond")

	g := MarsGrid{grid: [100][100]gridCell{},}

	log.Println(g.grid[0])
}

//Mars Grid cell
type gridCell struct{
	lifeChance int
	occupied bool

}

// MarsGrid represents a grid of some of the surface
// of Mars. It may be used concurrently by different
// goroutines.
type MarsGrid struct{
	mu sync.Mutex
	grid [100][100]gridCell
}