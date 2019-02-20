package main

import "fmt"

func main() {
	// Slice of int.
	// If have len [5]int, will be a array.
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8} // Composite literal.
	fmt.Printf("Capacity: %v\n", cap(s1))
	fmt.Printf("Length: %v\n", len(s1))
	fmt.Printf("s1: %v\n", s1)
}
