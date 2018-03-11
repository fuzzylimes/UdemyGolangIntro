package main

import (
	"fmt"
)

func zero(x int) int {
	x = 0
	return x
}

func main() {
	x := 5
	x = zero(x)
	fmt.Println(x)
}
