package main

import (
	"log"
)

func main(){

	h := Human{}
	h.safeToWalk = true
	log.Println(h)
}

type Human struct{
	Legs
}

type Legs struct{
	safeToWalk bool
}