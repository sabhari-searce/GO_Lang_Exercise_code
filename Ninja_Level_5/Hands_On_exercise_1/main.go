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

}
