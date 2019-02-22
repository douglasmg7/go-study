package main

import "fmt"

func main() {
	fmt.Println(completeName("Lucas", "Gert"))
}

func completeName(firstName, lastName string) string {
	return firstName + " " + lastName
}
