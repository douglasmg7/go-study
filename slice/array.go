package main

import (
	"fmt"
)

// Array.
// Documentation said, use slice instead.
// Len is part of type.
// If no len []int, will be a slice.
var a [5]int

func main() {
	a[1] = 2
	fmt.Println(a)
	fmt.Println(len(a))
}
