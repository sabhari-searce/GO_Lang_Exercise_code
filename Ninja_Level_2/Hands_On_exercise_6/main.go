package main

import "fmt"

const (
	first  = 2022 + iota
	second = 2022 + iota
	third  = 2022 + iota
	final  = 2022 + iota
)

func main() {

	fmt.Println("The four Years are ", first, second, third, final)
}
