package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: square root of a negative number: %v", f)
	}
	return 42, nil
}
