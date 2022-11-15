package main

import "fmt"

func main() {
	my_data := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, data := range my_data {
		fmt.Println("Printing data about ", data[0], data[1])
		for _, value := range data {
			fmt.Printf("\t%v\n", value)
		}
	}
}
