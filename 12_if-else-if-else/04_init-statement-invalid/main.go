package main

import (
	"fmt"
)

func main() {
	b := true
	if food := "Chocolate"; b { // Able to initialize a variable before doing the evalutation
		fmt.Println(food)
	}
	// fmt.Println(food)
	// The above falls out the scope where food is available, so it can't be used
}
