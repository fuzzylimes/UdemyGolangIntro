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
		ErrNorgateMath := fmt.Errorf("norgate math: square root of a negative number: %v", f)
		return 0, ErrNorgateMath
	}
	return 42, nil
}
