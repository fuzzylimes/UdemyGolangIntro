package main

import (
	"fmt"
)

func main() {
	var x rune = 'a' // alias for int32
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))
}
