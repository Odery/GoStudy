package main

import "fmt"

type coordinates struct{
	x,y float64
}

func (c *coordinates)moveUp(v float64){
	c.x += v
}
func (c *coordinates)moveDown(v float64){
	c.x -= v
}
func (c *coordinates)moveLeft(v float64){
	c.y -= v
}
func (c *coordinates)moveRight(v float64){
	c.y += v
}
type turtle struct{
	name string
	age int
	coordinates
}

func (t turtle) String() string{
	return fmt.Sprintf("%s is %d years old, coordinates: %+v", t.name, t.age, t.coordinates)
}
func main(){
	berry := turtle{name: "Berry",age: 213}

	fmt.Println(berry)

	berry.moveDown(23)
	berry.moveRight(10)
	berry.moveDown(0.3)

	fmt.Println(berry)
}