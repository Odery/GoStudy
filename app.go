package main

import (
	"log"
	"time"
)

func main(){
	log.Print(time.Millisecond)

	log.Print((250 * time.Millisecond))

	log.Print((250* time.Millisecond + 500 * time.Millisecond))

}