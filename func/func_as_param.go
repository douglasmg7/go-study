package main

import "fmt"

func main() {
	call(func(x int) { fmt.Println(x) }, 45)
}

func call(fn func(int), x int) {
	fn(x)
}
