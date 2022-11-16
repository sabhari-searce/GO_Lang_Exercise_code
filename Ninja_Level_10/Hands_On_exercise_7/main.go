package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go send(c)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}

	//for v := range c {
	//	fmt.Println(v)
	//}
}

func send(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	//close(c)
}
