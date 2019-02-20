package main

import "fmt"

func main() {
  // Slice of int.
  s1 := []int{1, 2, 3, 4} // Composite literal.
  s2 := []int{5, 6}
  // Create a new slice with appended arrays.
  fmt.Printf("s1: %v\n", s1)
  fmt.Printf("Capacity: %v\n", cap(s1))
  fmt.Printf("Length: %v\n\n", len(s1))
  fmt.Printf("s2: %v\n\n", s2)
  fmt.Printf("Capacity: %v\n", cap(s2))
  fmt.Printf("Length: %v\n\n", len(s2))

  // New underline array are created.
  // Append to s2 to s1.
  // Bad use, must be, "s1 := append(s1, s2...)".
  s3 := append(s1, s2...)
  s3[0] = 20
  fmt.Printf("s3: %v\n", s3)
  fmt.Printf("Capacity: %v\n", cap(s3))
  fmt.Printf("Length: %v\n\n", len(s3))
}
