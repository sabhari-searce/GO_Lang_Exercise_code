package main

import "fmt"

func main() {
	fmt.Println(evenNo(add, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func add(list_inp []int) int {
	var tot int
	for _, v := range list_inp {
		tot += v
	}

	return tot
}

func evenNo(f func(xi []int) int, list_inp ...int) int {
	var res int
	list := []int{}

	for _, elem := range list_inp {
		if elem%2 == 0 {
			list = append(list, elem)
		}
	}
	res = f(list)

	return res
}
