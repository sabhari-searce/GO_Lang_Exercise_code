package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("The number 42 in decimal is %d\n", x)
	fmt.Printf("The number 42 in binary is %b\n", x)
	fmt.Printf("The number 42 in hexa decimal is %x\n", x)

	x = x << 1
	fmt.Printf("The number 42 after shifting bits in decimal is %d\n", x)
	fmt.Printf("The number 42 after shifting bits in binary is %b\n", x)
	fmt.Printf("The number 42 after shifting bits in hexa decimal is %x\n", x)

}
