package main

import (
	"fmt"
	"net/url"
)

func main(){
	fmt.Println()
	_,err := url.Parse("https://a b.com/")
	if err != nil{
		fmt.Printf("%#v\n", err)
		if e, ok := err.(*url.Error); ok{
			fmt.Printf("%v, for operation %v for url: %v\n", e.Err, e.Op, e.URL)
		}
	}
}