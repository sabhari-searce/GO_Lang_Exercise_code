package main

import (
	"fmt"
	"math/rand"
)

func main() {
	my_slice := [10]int{}
	for i := 0; i < 10; i++ {
		my_slice[i] = rand.Intn(100)
	}
	for i, v := range my_slice {
		fmt.Println("the ", i, "th element is ", v)
	}
	fmt.Printf("The type of slice is %T\n", my_slice)
}
