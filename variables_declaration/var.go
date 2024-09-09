package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var i float64 = 4.98
	j := 5 + int(i)
	var s string = `Hello
World`

	// Need be initialized.
	const myConst string = "asdf"

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(s)

	// 1.
	fmt.Println(len("a"))
	// 2. It use two bytes to represent this character.kj
	fmt.Println(len("ã"))

	// Correct way to get the string size.
	fmt.Println(utf8.RuneCountInString("ãbc"))
	fmt.Println(utf8.RuneCountInString("abc"))
}