package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println("The number ", i, " when divided by 4 leaves ", i%4)
	}
}
