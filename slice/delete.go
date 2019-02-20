package main

import "fmt"

func main() {
  // Slice of int.
  s1 := []int{1, 2, 3, 4, 5, 6, 7, 8} // Composite literal.
  fmt.Printf("s1: %v\n", s1)
  fmt.Printf("Capacity: %v\n", cap(s1))
  fmt.Printf("Length: %v\n\n", len(s1))

  // Delete the third item.
  s2 := append(s1[:2], s1[3:]...)
  // Still using the same underline array, so s1 changed!.
  fmt.Printf("s1: %v\n", s1)
  fmt.Printf("Capacity: %v\n", cap(s1))
  fmt.Printf("Length: %v\n\n", len(s1))
  fmt.Printf("s2: %v\n", s2)
  fmt.Printf("Capacity: %v\n", cap(s2))
  fmt.Printf("Length: %v\n\n", len(s2))

}
