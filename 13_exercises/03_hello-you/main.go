package main

import (
	"fmt"
)

func main() {

	var you string

	fmt.Print("What is your name? ")
	fmt.Scan(&you)
	fmt.Println("Hello,", you)
}
