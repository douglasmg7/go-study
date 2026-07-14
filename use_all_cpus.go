package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Number of cpus: %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU()) // Uses all available CPU cores
}