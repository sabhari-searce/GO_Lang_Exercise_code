package main

import "fmt"

func main() {
	fmt.Println("please enter a value")
	func() {
		var inp int
		fmt.Scan(&inp)
	}()
	fmt.Println(inp)
}
