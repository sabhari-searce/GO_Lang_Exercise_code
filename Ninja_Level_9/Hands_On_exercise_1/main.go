package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("End of main")

	wg.Wait()
}

func foo() {
	fmt.Println("Hi from foo")
	wg.Done()
}

func bar() {
	fmt.Println("Hi from bar")
	wg.Done()
}
