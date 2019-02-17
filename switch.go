package main

import "fmt"

func main() {
	switch {
	case 2 == 2:
		fmt.Println("Will print")
	case 2 > 1:
		fmt.Println("Will not print")
	}

	// Break is the default.
	// Use fallthrough to cancel de break.
	food := "banana"
	switch food {
	case "orange", "apple", "banana":
		fmt.Println("A fruit")
		fallthrough
	case "dragonfruit":
		fmt.Println("A diferent fruit")
	default:
		fmt.Println("Not a fruit")
	}
}
