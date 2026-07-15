package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dx)
	for i := range s {
		// fmt.Println(i)
		s[i] = make([]uint8, dy)
		for j := range s[i] {
			s[i][j] = uint8(i * j)
		}
	}
	// fmt.Printf("s: %T\n", s)
	// fmt.Printf("s len: %v\n", len(s))
	// fmt.Printf("s capacity: %v\n", cap(s))

	// fmt.Printf("s: %T\n", s[0])
	// fmt.Printf("s len: %v\n", len(s[0]))
	// fmt.Printf("s capacity: %v\n", cap(s[0]))

	// for i, _ := range s {
	// for j, val_j := range s[i] {
	// fmt.Printf("[%d, %d] = [%d]\n", i, j, val_j)
	// }
	// }
	return s
}

func main() {
	pic.Show(Pic)
	// Pic(4, 5)
}