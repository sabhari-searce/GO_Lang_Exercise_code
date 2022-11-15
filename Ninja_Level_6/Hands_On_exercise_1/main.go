package main

import (
	"fmt"
)

func foo() int {
	var input int
	fmt.Println("please Enter a number")
	fmt.Scan(&input)

	return input
}

func bar(num int) (int, string) {
	return num, fmt.Sprintln("The input number is returned")

}

func main() {
	inp := foo()

	fmt.Println(bar(inp))

}
