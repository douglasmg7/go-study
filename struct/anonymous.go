package main

import "fmt"

func main() {
	p := struct {
		name string
		age  int
	}{
		name: "Lucas",
		age:  3,
	}
	fmt.Println(p)
}
