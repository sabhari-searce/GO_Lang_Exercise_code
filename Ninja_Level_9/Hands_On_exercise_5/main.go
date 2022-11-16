package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		//fmt.Println("The no of cpu and no of go routines are ", runtime.NumCPU(), runtime.NumGoroutine())
		go incrementor()
	}
	wg.Wait()

}

func incrementor() {
	atomic.AddInt64(&counter, 1)
	fmt.Println("the value of counter is ", counter)
	wg.Done()
}
