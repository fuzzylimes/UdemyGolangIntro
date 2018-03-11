package main

import (
	"fmt"
	"os"
)

func evenodd(x *int) string {
	if *x%2 > 0 {
		return "Odd"
	}
	return "Even"
}

func main() {
	var mynum int
	fmt.Print("Enter an int: ")
	_, err := fmt.Scanf("%d", &mynum)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(mynum, "is", evenodd(&mynum))

}
