package main

import "fmt"

func main() {
	f := invdSum

	odd, even := f(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("the odd and even sum is ", odd, even)

}

func invdSum(list_inp ...int) (int, int) {
	var oddSum, evenSum int

	for _, v := range list_inp {
		if v%2 == 0 {
			evenSum += v
		} else {
			oddSum += v
		}
	}
	return oddSum, evenSum

}
