package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	key, err := GenerateRandomBytes(64) // 64 bytes = 512 bits
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated key (base64 encoded): %s\n", base64.StdEncoding.EncodeToString(key))
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random number generator fails to function correctly.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}