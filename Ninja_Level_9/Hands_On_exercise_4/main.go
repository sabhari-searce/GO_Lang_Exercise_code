package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementor int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		fmt.Println("CPU and Go routine ", runtime.NumCPU(), runtime.NumGoroutine())
		go increment()
	}
	wg.Wait()

}

func increment() {
	mu.Lock()
	temp := incrementor
	//runtime.Gosched()
	temp++
	incrementor = temp
	fmt.Println("The value of incrementor is ", incrementor)
	wg.Done()
	mu.Unlock()
}
