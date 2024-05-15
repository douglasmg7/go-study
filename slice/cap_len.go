package main

import "fmt"

func main() {
	ar := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("ar: %v\n", ar)
	fmt.Printf("ar len: %d\t", len(ar))
	fmt.Printf("ar cap: %d\n\n", cap(ar))

	sl1 := ar[1:]
	sl1[2] = 22
	fmt.Printf("sl1: %v\n", sl1)
	fmt.Printf("sl1 len: %d\t", len(sl1))
	fmt.Printf("sl1 cap: %d\n\n", cap(sl1))

	sl2 := ar[:9]
	fmt.Printf("sl2: %v\n", sl2)
	fmt.Printf("sl2 len: %d\t", len(sl2))
	fmt.Printf("sl2 cap: %d\n\n", cap(sl2))

	sl2 = append(sl2, 11)
	sl2 = append(sl2, 12)
	sl2[5] = 55
	fmt.Printf("sl2: %v\n", sl2)
	fmt.Printf("sl2 len: %d\t", len(sl2))
	fmt.Printf("sl2 cap: %d\n\n", cap(sl2))

	fmt.Printf("ar: %v\n", ar)
	fmt.Printf("ar len: %d\t", len(ar))
	fmt.Printf("ar cap: %d\n\n", cap(ar))

	fmt.Printf("sl1: %v\n", sl1)
	fmt.Printf("sl1 len: %d\t", len(sl1))
	fmt.Printf("sl1 cap: %d\n\n", cap(sl1))
}