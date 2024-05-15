package main

import "fmt"

func main() {
	m := map[string]int{}
	m["a"] = 1
	fmt.Printf("m: %v\n", m)
	fmt.Printf("have the key: %v\n", m["a"])
	fmt.Printf("No have the key: %v\n\n", m["b"])

	m2 := map[string]int{
		"b": 2,
		"c": 3,
	}
	fmt.Printf("m2: %v\n", m2)

	// Using make and giving a capacity.
	m3 := make(map[string]int, 2)
	m3["ana"] = 3
	m3["pricila"] = 7
	fmt.Printf("m3: %v\n", m3)
}
