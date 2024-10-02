package main

import "fmt"

func main() {
	car := "Cobalt"
	i := interface{}(car)
	fmt.Printf("i: %T\n", i)

	fmt.Println(i.(string))

	val, ok := i.(string)
	if ok {
		fmt.Println("Is a string")
		fmt.Println(val)
	}
}