package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"sync"
	"time"
)

const(
	gridHeight = 10
	gridWidth = 25
)

func main() {
	earth := NewEarth()
	grid := NewMarsGrid()
	
	StartRover("Curiosity", grid, image.Point{X:1, Y:1}, earth.in)
	StartRover("Oportunity", grid, image.Point{X:3, Y:3}, earth.in)
	StartRover("Calisto", grid, image.Point{X:6, Y:6}, earth.in)
	StartRover("Huan", grid, image.Point{X:8, Y:9}, earth.in)

	for {
		report := <- earth.out
		for _, s := range report {
			log.Println(s)
        }
	}
}
 
type Cell struct{
	mu sync.Mutex
	chanceOfLife int
	reported bool
	Occupier
}

type MarsGrid struct{
	grid [gridHeight][gridWidth]Cell
}

func NewMarsGrid() *MarsGrid{
	m := MarsGrid{}

	for i := range m.grid{
		for j := range m.grid[i]{
			m.grid[i][j] = Cell{
				chanceOfLife: rand.Intn(1000) + 1,
				Occupier: Occupier{
					position: image.Point{i,j},
					mg: &m,
				},
			}
		}
	}

	return &m
}

func (m *MarsGrid) Occupy(p image.Point) *Occupier{
	cell := &m.grid[p.X][p.Y]

	if cell.Occupier.occupied == false && inBounds(p){
		cell.mu.Lock()
		defer cell.mu.Unlock()

		oc := cell.Occupier
		oc.occupied = true
		return &oc
	}else{
		return nil
	}
}

func inBounds(p image.Point) bool{
	return p.X >= 0 && p.X < gridHeight && p.Y >= 0 && p.Y < gridWidth
}

type Occupier struct{
	position image.Point
	occupied bool
	mg *MarsGrid
}

func (o *Occupier) Move(p image.Point) bool{
	var targetCell *Cell
	if inBounds(p){
		targetCell = &o.mg.grid[p.X][p.Y]
	}else{
		return false
	}
	if targetCell.occupied == false{
		targetCell.mu.Lock()
		defer targetCell.mu.Unlock()

		currentCell := &o.mg.grid[o.position.X][o.position.Y]
		currentCell.mu.Lock()
		defer currentCell.mu.Unlock()

		currentCell.occupied = false
		targetCell.occupied = true

		o.position = p
		
        return true
	}
	return false
}

type Rover struct{
	mu sync.Mutex
	name string
	earth chan string
	mg *MarsGrid
	*Occupier
}


func StartRover(name string, mg *MarsGrid, startPoint image.Point, ch chan string){
	r := Rover{
		mg: mg,
		name: name,
		earth: ch,
		Occupier: mg.Occupy(startPoint),
	}
	go r.Drive()
}

func (r *Rover) Drive(){
	for {
		next := time.After(250 * time.Millisecond)
		select{
		case <- next:
			if ok, life := r.cellHasLife(); ok{
				r.earth <- fmt.Sprintf("Rover: %v, found life in cell, life value is: %v. Cell coordinates: %v", r.name, life, r.position)
			}
			r.seekLife()
		}
	}
}

func (r *Rover) seekLife(){
	r.mu.Lock()
	defer r.mu.Unlock()
	
	for moved := false; !moved; {
		switch c := rand.Intn(4); c {
			case 0:
                moved = r.goLeft()
			case 1:
                moved = r.goRight()
            case 2:
                moved = r.goUp()
            case 3:
                moved = r.goDown()
		}
	}
}

func (r *Rover) goLeft() bool{
	return r.Move(image.Point{X: r.Occupier.position.X, Y: (r.Occupier.position.Y - 1)})
}

func (r *Rover) goRight() bool{
	return r.Move(image.Point{X: r.Occupier.position.X, Y: (r.Occupier.position.Y + 1)})
}

func (r *Rover) goDown() bool{
	return r.Move(image.Point{X: (r.Occupier.position.X + 1), Y: r.Occupier.position.Y})
}

func (r *Rover) goUp() bool{
	return r.Move(image.Point{X: (r.Occupier.position.X - 1), Y: r.Occupier.position.Y})
}

func (r *Rover) cellHasLife() (bool, int){
	cell := &r.mg.grid[r.position.X][r.position.Y]
	if !cell.reported {
		cell.reported = true
		return cell.chanceOfLife > 900, cell.chanceOfLife
	}else{
		return false, 0
	}
}

type Earth struct{
	in chan string
	out chan []string
}

func (e *Earth) listenToRovers() {
	report := make([]string, 0)
	for {
		next := time.After(1 * time.Second)
        select {
        case msg := <- e.in:
            report = append(report, msg)
		case <- next:
			if len(report) > 0 {
				log.Println("Data received from rovers")
				e.out <- report
				report = make([]string, 0)
        	}else{
				log.Println("No data received from rovers")
			}
    	}
	}
}

func NewEarth() *Earth {
	earth := Earth{
		in: make(chan string),
        out: make(chan []string),
	}

	go earth.listenToRovers()

	return &earth
}
