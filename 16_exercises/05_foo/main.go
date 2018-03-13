package main

import (
	"fmt"
)

func foo(si ...int) {
	i := 0
	for _, v := range si {
		if v > i {
			i = v
		}
	}
	fmt.Println(i)
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
