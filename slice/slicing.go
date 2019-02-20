package main

import "fmt"

func main() {
  // Slice of int.
  s1 := []int{1, 2, 3, 4, 5, 6, 7, 8} // Composite literal.

  // Slicing, using the same underline array.
  s2 := s1[1:]
  s2[0] = 20
  fmt.Printf("s1: %v\n", s1)
  fmt.Printf("Capacity: %v\n", cap(s1))
  fmt.Printf("Length: %v\n\n", len(s1))

  fmt.Printf("s2: %v\n", s2)
  fmt.Printf("Capacity: %v\n", cap(s2))
  fmt.Printf("Length: %v\n", len(s2))

}
