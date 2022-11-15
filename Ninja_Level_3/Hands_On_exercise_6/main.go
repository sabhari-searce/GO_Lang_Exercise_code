package main

import "fmt"

func main() {
	var number int
	fmt.Println("Please enter a number")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println(number, " is a even number")
	}
}
