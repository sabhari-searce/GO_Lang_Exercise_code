package main

import "fmt"

func main() {
	//shorthand assignment to variable
	x := 42
	y := "James Bond"
	z := true

	//all variables in single print statement
	fmt.Println("The values of the variables are ", x, y, z)

	//using multiple print statements
	fmt.Printf("The value of variable x is %v and its type is %T\n", x, x)
	fmt.Printf("The value of variable y is %v and its type is %T\n", y, y)
	fmt.Printf("The value of variable z is %v and its type is %T\n", z, z)
}
