package main

import (
	"fmt"
)

func sum(c chan int) {
	x := <-c
	y := <-c
	c <- x + y
}

func main() {
	c := make(chan int, 1)
	go sum(c)
	c <- 10
	c <- 8
	result := <-c
	fmt.Println(result)
}