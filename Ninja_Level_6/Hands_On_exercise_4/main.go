package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi my name is ", p.first, p.last, " and i'm ", p.age, " years old!!")
}
func main() {

	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p.speak()
}
