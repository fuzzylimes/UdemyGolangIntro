package main

import (
	"fmt"
)

func half(num int) (int, bool) {
	h := num / 2
	if a := num % 2; a != 0 {
		return h, false
	}
	return h, true
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
