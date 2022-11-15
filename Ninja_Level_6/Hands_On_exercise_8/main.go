package main

import "fmt"

func main() {
	f := choosing("add")
	fmt.Println(f(5, 6))

	f = choosing("subtract")
	fmt.Println(f(6, 5))

}

func choosing(para string) func(a int, b int) int {
	switch para {
	case "add":
		return add
	case "subtract":
		return subtract
	}
	return add
}

func subtract(a int, b int) int {
	return a - b

}

func add(a, b int) int {
	return a + b
}
