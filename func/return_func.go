package main

import "fmt"

func main() {
	fmt.Println(foo())
	x := bar()
	fmt.Printf("x = %+T\n", x)
	fmt.Println(x())
	fmt.Println(bar()())
}

func foo() string {
	return "Hello world"
}

func bar() func() int {
	return func() int {
		return 451
	}
}
