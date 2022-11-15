package main

import "fmt"

func main() {
	Birth_year := 2000
	for true {
		if Birth_year > 2022 {
			break
		}
		fmt.Println(Birth_year)
		Birth_year++
	}
}
