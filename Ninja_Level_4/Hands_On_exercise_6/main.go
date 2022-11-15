package main

import "fmt"

func main() {
	states := make([]string, 50, 50)
	fmt.Println("The length of array is ", len(states))
	fmt.Println("The capacity of array is ", cap(states))
	fmt.Println("the contents are ", states)
	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println("The length of array is ", len(states))
	fmt.Println("The capacity of array is ", cap(states))
	for i, v := range states {
		fmt.Println(i, v)
	}

	new_state := make([]string, 50, 50)
	for i := 0; i < len(new_state); i++ {
		new_state[i] = states[50+i]
	}
	fmt.Println("The length of array is ", len(new_state))
	fmt.Println("The capacity of array is ", cap(new_state))
	for i, v := range new_state {
		fmt.Println(i, v)
	}
}
