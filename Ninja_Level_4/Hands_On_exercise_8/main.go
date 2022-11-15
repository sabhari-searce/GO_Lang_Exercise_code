package main

import "fmt"

func main() {
	my_map := map[string][]string{}

	my_map["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	my_map["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	my_map["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for k, v := range my_map {
		fmt.Println("The name is ", k)
		for i, l := range v {
			fmt.Println("\tthe index is ", i, " the value is ", l)
		}
	}

}
