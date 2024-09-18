package main

import "fmt"

func main() {
	type dir string
	s1 := dir("asdf")
	i1 := float64(2)
	fmt.Printf("s1 type: %T\n", s1)
	fmt.Printf("i1 type: %T\n", i1)
}