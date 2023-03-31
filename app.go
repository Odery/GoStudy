package main

import "fmt"

type cat struct{
	name string
}

func (c cat) voice(s string, count int){
	for i :=0; i < count; i++{
		fmt.Println(s)
	}
}

type dog struct{
	name string
}

func (d dog) voice(s string){
	fmt.Println("Woof! ",s)
}

type zoo struct{
	cat
	dog
}
func main() {
	z := zoo{cat{name: "Bob"},dog{name: "Jhon"}}

	z.voice("Barq", 2)
}