package main

import "fmt"

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	country string
}

func main() {
	p := person{
		name: "Lucas",
		age:  3,
	}

	sa := secretAgent{
		person: person{
			name: "Luck",
			age:  23,
		},
		country: "Brasil",
	}

	fmt.Println(p)
	fmt.Println(sa)
}
