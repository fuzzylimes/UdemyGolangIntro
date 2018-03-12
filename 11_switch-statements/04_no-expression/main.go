package main

import (
	"fmt"
)

func main() {
	myFriendsName := "Marcus"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("My friend has a name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Hello Tim")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Hello Marcus/Medhi")
	default:
		fmt.Println("Nothing match; this is the default")
	}
}
