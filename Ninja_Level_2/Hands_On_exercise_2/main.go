package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 != 42)
	c := (42 >= 40)
	d := (42 <= 13)
	e := (1 < 2)
	f := (1 > 2)

	fmt.Println(a, b, c, d, e, f)
}
