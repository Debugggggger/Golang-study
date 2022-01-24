package main

import (
	"fmt"
	"os"
)

func main(){
	for index,env := range os.Environ() {
		fmt.Println(index,env)
	}
}