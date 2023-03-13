/*
Cipher the Spanish message “Hola, Estación Espacial Internacional!” with ROT13. 
Use the range keyword. Now when you use ROT13 on Spanish text,
characters with accents are preserved
*/

package main

import(
	"fmt"
)

func main(){
	var message = "Hola, Estación Espacial Internacional!"
	var encoded string

	for _,c := range(message){
		if c >= 'A' && c <= 'z'{
			c += 13
			if c > 'z'{
				c -= 26
			}
			encoded += fmt.Sprintf("%c",c)
		}else{
			encoded += fmt.Sprintf("%c",c)
		}
	}
	fmt.Printf("Encoded phrase: %v", encoded)
}