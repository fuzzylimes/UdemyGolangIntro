package main

import (
	"fmt"
)

func greatest(si ...int) int {
	var i int
	for _, v := range si {
		if v > i {
			i = v
		}
	}
	return i
}

func main() {
	n := greatest(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}
