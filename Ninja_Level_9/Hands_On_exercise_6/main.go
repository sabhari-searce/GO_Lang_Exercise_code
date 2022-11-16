package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("The Os and architecture are ", runtime.GOOS, runtime.GOARCH)
}
