package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("The Value of the variables x, y, z are %v, %v, %v", x, y, z)
	fmt.Println(s)

}
