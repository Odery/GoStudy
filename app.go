package main

import (
	"fmt"
	"bytes"
)

func main() {
	input := "Brandbury is at: 130.00º20.00'0.00\" H, 27.00º13.00'2.20\" E"

	// Convert the string to a slice of runes
	runes := []rune(input)

	// Create a bytes.Buffer to store the converted runes as bytes
	var buf bytes.Buffer

	// Iterate through the runes and write each rune as a byte to the buffer
	for _, r := range runes {
		buf.WriteRune(r)
	}

	// Get the byte array from the buffer
	byteArray := buf.Bytes()

	// Print the byte array
	fmt.Println(byteArray)

	// Convert the byte array back to a string and print it
	output := string(byteArray)
	fmt.Println(output)
}