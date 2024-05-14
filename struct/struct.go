package main

import "fmt"

// Define a struct.
type person struct {
	name string
	age  int
}

// Define a struct.
type secretAgent struct {
	person
	country string
}

func main() {
	// Long notation.
	// Declaration + assign.
	p := person{
		name: "Lucas",
		age:  3,
	}

	// Short notation.
	// Using the same order as the struct definition.
	p3 := person{
		"Lilian",
		24,
	}

	// Field that use a struct.
	sa := secretAgent{
		person: person{
			name: "Luck",
			age:  23,
		},
		country: "Brasil",
	}

	// Declaration.
	var p2 person
	// Assign
	p2 = person{
		name: "Júlia",
		age:  32,
	}

	// Anonymous declaration.
	persion_anonymous := struct {
		name string
		age  int
	}{
		name: "Lucas",
		age:  3,
	}

	// Empty person.
	empty_person := person{}

	// Empty field.
	empty_person_field := person{
		name: "Lulu",
	}

	fmt.Println(p)
	fmt.Println(sa)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(persion_anonymous)
	fmt.Println(empty_person)
	fmt.Println(empty_person_field)

}
