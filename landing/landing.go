package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct{
	Name string `json:"name"`
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

func main(){
	landings := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Lon: 137.4417},
 		{Name: "Columbia Memorial Station", Lat: -14.5684, Lon: 175.472636},
 		{Name: "Challenger Memorial Station", Lat: -1.9462, Lon: 354.4734},
	}
	mar, err := json.MarshalIndent(landings," ","  ")
	exitOnErr(err)

	fmt.Println(string(mar))
}

func exitOnErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}