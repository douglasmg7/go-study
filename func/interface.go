package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	fn string
	ln string
}

type agent struct {
	name string
}

func main() {
	p := person{
		fn: "Lucas",
		ln: "Zair",
	}
	a := agent{
		name: "---",
	}
	print(p)
	print(a)
}

func (p person) speak() {
	fmt.Println(p)
}

func (a agent) speak() {
	fmt.Println(a)
}

func print(h human) {
	h.speak()
	switch h.(type) {
	case person:
		fmt.Println("He is a person.")
	case agent:
		fmt.Println("He is a agent.")
	default:
		fmt.Println("Do not know.")
	}
}
