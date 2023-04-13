package main

import "fmt"

type item string

type character struct{
	name string
	item *item
}

func (c *character) pickup(i *item){
	c.item = i
	fmt.Printf("%s has picked up the %s\n", c.name, *i)
}

func (c *character)giveItem(b *character){
	if c.item != nil{
		b.item = c.item
		c.item = nil
		fmt.Printf("%s has given %s to the %s\n", c.name, *b.item, b.name)
	}else{
		fmt.Println("No item in inventory")
	}
}

func main(){
	arthur := character{name: "Arthur"}
	lancelot := character{name: "Lancelot"}
	sword := item("Long Sword")

	arthur.pickup(&sword)

	arthur.giveItem(&lancelot)
}