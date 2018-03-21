package main

import (
	"fmt"
)

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v\n", myAge, myAge)

	var yourAge int
	myAge = 29
	fmt.Printf("%T %v\n", yourAge, yourAge)

	fmt.Println(int(myAge) + yourAge)
	fmt.Println(myAge + foo(yourAge))
}
