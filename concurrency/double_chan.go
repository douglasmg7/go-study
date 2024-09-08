package main

import "fmt"

func sum(ci chan int, co chan int) {
	for {
		x := <-ci
		y := <-ci
		fmt.Println(x, "+", y, "=", x+y)
		co <- x + y
	}
}

func main() {
	ci := make(chan int, 6)
	co := make(chan int, 3)
	go sum(ci, co)

	ci <- 1
	ci <- 2
	ci <- 3
	ci <- 4
	ci <- 5
	ci <- 6
	result := <-co
	fmt.Println(result)
	result = <-co
	fmt.Println(result)
	result = <-co
	fmt.Println(result)
}