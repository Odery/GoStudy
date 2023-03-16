/*
Write a program to decipher the ciphered text. To keep it simple, all
characters are uppercase English letters (65 till 90) for both the text and keyword:
cipherText := "CSOITEUIWUIZNSROCNKFD"
keyword := "GOLANG"
 The strings.Repeat function may come in handy. Give it a try, but also complete
this exercise without importing any packages other than fmt to print the deciphered message.
 Try this exercise using range in a loop and again without it. Remember that the
range keyword splits a string into runes, whereas an index like keyword[0] results in
a byte.
#TIP# You can only perform operations on values of the same type, but you can convert one
type to the other (string, byte, rune).
 To wrap around at the edges of the alphabet, the Caesar cipher exercise made use of
a comparison. Solve this exercise without any if statements by using modulus (%).

Also make a func to cipher text.
Bonus: rather than write your plain text message in uppercase letters with no spaces,
use the strings.Replace and strings.ToUpper functions to remove spaces and uppercase the
string before you cipher it.
Once you’ve ciphered a plain text message, check your work by deciphering the
ciphered text with the same keyword
*/

package main

import (
	"fmt"
	"strings"
)

func main(){
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	plainText := "Hello my name is Joe"
	keyword := "GOLANG"

	fmt.Printf("Deciphered given text: %v\n",decipherNew(cipherText, keyword))

	fmt.Printf("Ciphered text: %v\n", cipher(plainText, keyword))
	fmt.Printf("Deciphered: %v\n", decipherNew(cipher(plainText,keyword),keyword))

}

// With range Version1
func decipher(cipherText string, keyword string) string{
	var deciphered string

	for i, c := range(cipherText){
		keyRune := rune(keyword[(i % 6)])
		decipheredChar := (c + ('A' - keyRune))
		deciphered += fmt.Sprintf("%c",(((decipheredChar - 64) % 26 + 25) % 26 + 65))
	}
	return deciphered
}

//Without range() and with explained formula to keep values between 'A' and 'Z' ('65' and '90')
//Formula to restrict value between value x and y is 'result = ((input_value - x) % (y - x + 1) + (y - x + 1)) % (y - x + 1) + x'
func decipherNew(cipherText string, keyword string) string{
	var deciphered string

	for i := 0; i< len(cipherText); i++{
		keyRune := keyword[i % 6]
		decipheredChar := (cipherText[i] + ('A' - keyRune))
		x,y := 65,90
		e := (y - x + 1)
		deciphered += fmt.Sprintf("%c",((int(decipheredChar) - x) % e + e) % e + x )
	}
	return deciphered
}

func cipher(plainText string, keyword string) string{
	var ciphered string
	plainText = strings.Replace(plainText, " ", "", -1)
	plainText = strings.ToUpper(plainText)

	for i, c := range(plainText){
		keyRune := rune(keyword[i % 6])
		cipheredChar := (c + (keyRune - 'A'))
		x,y := 65,90
		e := (y - x + 1)
		ciphered += fmt.Sprintf("%c",((int(cipheredChar) - x) % e + e) % e + x )
	}
	return ciphered
}