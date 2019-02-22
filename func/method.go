package main

import "fmt"

type person struct {
	fn string
	ln string
}

func (p person) speak() {
	fmt.Println(p)
}

func main() {
	p := person{
		fn: "Lucas",
		ln: "Zair",
	}
	p.speak()
}
