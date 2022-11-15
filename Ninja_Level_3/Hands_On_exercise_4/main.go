package main

import "fmt"

func main() {
	Year := 2000
	for {
		if Year > 2022 {
			break
		}
		fmt.Println(Year)
		Year++
	}
}
