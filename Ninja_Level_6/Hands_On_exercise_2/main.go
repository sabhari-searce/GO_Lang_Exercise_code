package main

import "fmt"

func foo(varpar ...int) int {
	var total int
	for _, v := range varpar {
		total += v
	}
	return total
}

func bar(inpList []int) int {
	var total int
	for _, v := range inpList {
		total += v
	}
	return total
}
func main() {
	fmt.Println(foo(1, 2, 3, 4, 5, 6, 7, 8))
	my_list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(bar(my_list))
}
