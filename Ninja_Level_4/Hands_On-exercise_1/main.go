package main

import "fmt"

func main() {
	my_array := [5]int{}
	my_array[0] = 54
	my_array[1] = 41
	my_array[2] = 90
	my_array[3] = 61
	my_array[4] = 9
	for i, v := range my_array {
		fmt.Println(i, "th emelemt is ", v)
	}
	fmt.Printf("The type of array is %T\n", my_array)
}
