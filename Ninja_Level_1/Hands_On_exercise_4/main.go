package main

import "fmt"

type my_very_own_type int

var x my_very_own_type

func main() {
	fmt.Printf("The value of variable x is %v\n", x)
	fmt.Printf("The type of variable x is %T\n", x)
	x = 42
	fmt.Printf("The value of variable x is %v\n", x)
	fmt.Printf("The type of variable x is %T\n", x)
}
