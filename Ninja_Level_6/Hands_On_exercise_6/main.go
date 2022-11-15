package main

import "fmt"

func main() {
	var start, end int
	func() {
		fmt.Println("Enter the start and end index")

		fmt.Scan(&start)
		fmt.Scan(&end)
	}()
	func(s int, e int) {
		fmt.Println("the numbers in the range are")
		for i := s; i <= e; i++ {
			fmt.Println(i)
		}
	}(start, end)
}
