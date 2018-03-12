package main

import (
	"fmt"
)

func main() {
	var small, large int

	fmt.Print("Pick a small number: ")
	fmt.Scan(&small)
	fmt.Print("Pick a large number: ")
	fmt.Scan(&large)
	fmt.Println("The remainder from", large, "divided by", small, "is", large%small)
}
