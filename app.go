package main

import (
	"fmt"
	"strings"
)

type talker interface{
	talk() string
}
type laser int

func (l laser) talk() string{
	return (strings.Repeat("pew", int(l)))
}
func shout(t talker){
	fmt.Println(strings.ToUpper(t.talk()))
}
func main() {
	la := laser(2)

	shout(&la)
}