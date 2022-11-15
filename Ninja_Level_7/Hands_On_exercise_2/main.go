package main

import "fmt"

type person struct {
	first   string
	last    string
	age     int
	address string
}

func changeMe(p *person) {
	(*p).address = "Bangalore"
}
func main() {
	p := person{
		first:   "James",
		last:    "Bond",
		age:     32,
		address: "America",
	}

	fmt.Println(p)

	changeMe(&p)
	fmt.Println(p)

}
