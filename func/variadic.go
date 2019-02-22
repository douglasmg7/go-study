package main

import "fmt"

func main() {
	moon(1, 2, 3, 4, 5, 6, 7)

	ii := []int{2, 3, 4, 5, 6, 7}
	moon(ii...)
	moon()
}

func moon(x ...int) {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Printf("len: %v, cap: %v\n", len(x), cap(x))
	fmt.Println(x[2])
	fmt.Println("")

}
