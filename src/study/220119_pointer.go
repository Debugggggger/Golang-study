package main

import "fmt"

func main() {
	var numPtr *int = new(int)
	
	fmt.Println(numPtr)

	x:=5
	zero(&x)
	fmt.Println(&x)
	fmt.Println(x)
}

func zero(xPtr *int){
	*xPtr = 0
}
