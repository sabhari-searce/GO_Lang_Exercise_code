package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementor int
var wg sync.WaitGroup

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		fmt.Println("CPU and Go routine ", runtime.NumCPU(), runtime.NumGoroutine())
		go increment()
	}
	wg.Wait()

}

func increment() {
	temp := incrementor
	runtime.Gosched()
	temp++
	incrementor = temp
	fmt.Println("The value of incrementor is ", incrementor)
	wg.Done()
}
