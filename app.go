package main

import "fmt"

type Person struct {
	name string
	belongings []int
}

func main() {
	bob := new(Person)

	for i:= 0; i < 10; i++ {
		bob.belongings = append(bob.belongings, i)
	}

	sl := make([]int, 5)

	copy(sl, bob.belongings[5:])
	fmt.Println(bob.name)
	fmt.Println(sl)
}