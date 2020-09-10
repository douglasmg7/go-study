package main

import "fmt"

func main() {
	s := "Hello"
	t := s
	s += " 1"
	fmt.Printf("s: %s\n", s)
	fmt.Printf("t: %s\n", t)
	t += " 2"
	fmt.Printf("t: %s\n", t)
}
