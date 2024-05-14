package main

import "fmt"

func main() {
	s := "Ula Ula"
	printAny(s)

	val := 3.0
	printIfInt(val)
}

// func printAny(i interface{}) {
func printAny(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Int value: %d\n", v)
	case string:
		fmt.Printf("String value: %s\n", v)
	default:
		fmt.Printf("I do not know the type\n")
	}
}

func printIfInt(i interface{}) {
	v, ok := i.(int)
	fmt.Printf("v: %d\n", v)
	if ok {
		fmt.Printf("It is a int value: %d\n", v)
		return
	} else {
		fmt.Println("Not a int value.")
	}
}
