package main

import "fmt"

func main() {
	var favSport string
	fmt.Println("please enter your favourite sport")
	fmt.Scan(&favSport)

	switch favSport {
	case "Cricket":
		fmt.Println("Cricket is an excellent sport")
	case "Football":
		fmt.Println("Lot of stamina to play football")
	case "Basketball":
		fmt.Println("Height matters")
	default:
		fmt.Println("I don't know that sport")

	}
}
