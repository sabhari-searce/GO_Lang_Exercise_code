package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favFlavours []string
}

func main() {
	p1 := person{
		firstName:   "James",
		lastName:    "Bond",
		favFlavours: []string{"vanilla", "choclate", "strawberry"},
	}

	p2 := person{
		firstName:   "Miss",
		lastName:    "MoneyPenny",
		favFlavours: []string{"Blueberry", "black current", "butter scotch"},
	}

	var persons = []person{p1, p2}
	//fmt.Println(persons)

	for _, v := range persons {
		fmt.Println("Their name is ", v.firstName, v.lastName)
		fmt.Println("Their favourite ice cream flavours are ")
		for _, value := range v.favFlavours {
			fmt.Printf("\t%v\n", value)
		}
	}

	my_map := map[string][]string{}
	for _, p := range persons {
		my_map[p.lastName] = p.favFlavours
	}

	for k, v := range my_map {
		fmt.Println("The name is ", k, " and the favourite flavours are ")
		for _, elem := range v {
			fmt.Printf("\t%v\n", elem)
		}
	}

}
