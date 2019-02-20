package main

import "fmt"

func main() {
	m := map[string]int{}
	m["a"] = 123
	m["d"] = 453
	v, ok := m["a"]
	fmt.Printf("For a - v: %v, ok: %v\n", v, ok)
	v, ok = m["b"]
	fmt.Printf("For b - v: %v, ok: %v\n", v, ok)

	// Test value.
	if _, ok := m["a"]; ok {
		fmt.Printf("Value exist")
	}

	// Range.
	for k, v := range m {
		fmt.Printf("\nKey: %v, value: %v", k, v)
	}
	fmt.Println()
}
