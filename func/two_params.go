package main

import "fmt"

func main() {
	a, b := invert("a", "b")
	fmt.Printf("a=%v, b=%v\n", a, b)
}

func invert(a, b string) (string, string) {
	return b, a
}
