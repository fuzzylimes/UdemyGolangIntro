package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println("a -", a)
	fmt.Println("a's address in memory is", &a)
	fmt.Printf("a's address in memory is %d\n", &a)
}
