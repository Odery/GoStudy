/*
Decipher the quote from Julius Caesar:
L fdph, L vdz, L frqtxhuhg.
						—Julius Caesar
Your program will need to shift uppercase and lowercase letters by –3. Remember that 
'a' becomes 'x', 'b' becomes 'y', and 'c' becomes 'z', and likewise for uppercase letters
*/
package main

import(
	"fmt"

)

func main(){
	var phrase = "L fdph, L vdz, L frqtxhuhg."
	var decoded string

	for _, c := range(phrase){
		if c >= 'A' && c <= 'z'{
			c -= 3
			decoded += fmt.Sprintf("%c",c)
		}else{
			decoded += fmt.Sprintf("%c",c)
		}
	}

	fmt.Printf("Decoded Caesar phrase: %v",decoded)
}