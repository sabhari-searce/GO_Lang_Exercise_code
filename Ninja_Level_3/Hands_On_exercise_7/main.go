package main

import "fmt"

func main() {
	var number int

	fmt.Println("Please enter a number between 0 and 100")
	fmt.Scan(&number)

	if number >= 0 && number < 10 {
		fmt.Println(number, " is a one digit number")
	} else if number >= 10 && number < 100 {
		fmt.Println(number, " is a two digit number")

	} else {
		fmt.Println("You had entered a larger number")
	}
}
