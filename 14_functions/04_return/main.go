package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname) // prints to a string rather than std out
}
