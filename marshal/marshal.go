package main

import (
	"encoding/json"
	"fmt"
	"bytes"
)

type coordinate struct{
	D,M,S float64
	H rune
}

func (c coordinate) MarshalJSON() ([]byte, error){
	var buf bytes.Buffer
	buf.WriteRune('{')

	i := 5
	for name, value := range(c.getData()){
		if i > 0{
			buf.WriteString(fmt.Sprintf("%#v: %#v,",name,value))
			i--
		}else{
			buf.WriteString(fmt.Sprintf("%#v: %#v",name,value))
		}
	}

	buf.WriteRune('}')
	return buf.Bytes(), nil
}

func (c coordinate) getData() map[string] interface{}{
	v := make(map[string]interface{}, 6)
	v["decimal"] = c.decimal()
	v["dms"] = fmt.Sprintf(`"%0.3fº%0.3f'%0.3f" '%c'"`,c.D, c.M, c.S, c.H)
	v["degrees"] = c.D
	v["minutes"] = c.M
	v["seconds"] = c.S
	v["hemisphere"] = string(c.H)

	return v
}
   
func (c coordinate) decimal() float64{
	sign := 1.0
	switch c.H{
	case 'S','W','s','w':
		sign = -1.0
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}


func main(){
	coor := coordinate{D:135, M: 54, H: 'E'}
	b,_ := json.MarshalIndent(coor, "","\t")
	fmt.Println(string(b))
}

