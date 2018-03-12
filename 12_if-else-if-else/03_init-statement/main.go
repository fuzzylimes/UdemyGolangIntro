package main

import (
	"fmt"
)

func main() {
	b := true
	if food := "Chocolate"; b { // Able to initialize a variable before doing the evalutation
		fmt.Println(food)
	}
}
