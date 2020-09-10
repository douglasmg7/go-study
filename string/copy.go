package main

import "fmt"

func main() {
	s := "Hello"
	newBuffer := make([]byte, len(s))
	copy(newBuffer, s)
	newBuffer[0] = 'q'
	t := string(newBuffer)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("t: %s\n", t)
	newBuffer[0] = 'r'
	fmt.Printf("t: %s\n", t)
}
