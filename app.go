package main

import (
	"fmt"
	"strings"
)

func main() {

	plainText := "Gooogle is bad"

	plainText = strings.Replace(plainText, " ", "", -1)
	plainText = strings.ToUpper(plainText)

	fmt.Println(plainText)
}
