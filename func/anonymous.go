package main

import "fmt"

func main() {
	func(x int) {
		fmt.Printf("Anonymous func with value: %v\n", x)
	}(4)
}
