package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle{
			doors:  2,
			colour: "Red",
		},
		true,
	}

	s1 := sedan{
		vehicle{
			doors:  4,
			colour: "White",
		},
		false,
	}

	//var vehicles = []vehicle{t1,s1}
	fmt.Println("The truck values are ", t1)
	fmt.Println("The sedan values are ", s1)

}
