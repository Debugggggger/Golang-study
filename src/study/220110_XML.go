package main

import (
	"encoding/xml"
	"fmt"
)

type Member struct {
	NAme string
	Age int
	Active bool
}

func main() {
	mem := Member {"Alex",10,true}

	xmlBytes, err := xml.Marshal(mem)
	if err != nil{
		panic(err)
	}

	xmlString := string(xmlBytes)
	fmt.Println(xmlString)
}