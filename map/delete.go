package main

import "fmt"

func main() {
	m := map[string]int{}
	m["a"] = 123
	m["d"] = 453

	fmt.Println(m)
	// Delete a value.
	delete(m, "a")
	fmt.Println(m)
}
