package main

import "fmt"

func main() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", append(s1, s2...))
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
}
