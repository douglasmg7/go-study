package main

import "fmt"

func main() {
  s1 := []int{}
  fmt.Printf("s1: %v\n", s1)
  fmt.Printf("Capacity: %v\n", cap(s1))
  fmt.Printf("Length: %v\n\n", len(s1))

  // Create a slice (type, len, cap).
  s1 = make([]int, 2, 4)
  fmt.Printf("s1: %v\n", s1)
  fmt.Printf("Capacity: %v\n", cap(s1))
  fmt.Printf("Length: %v\n\n", len(s1))
  s2 := s1
  fmt.Printf("s2: %v\n", s2)
  fmt.Printf("Capacity: %v\n", cap(s2))
  fmt.Printf("Length: %v\n\n", len(s2))

  s1 = append(s1, 3, 4)
  fmt.Printf("s1: %v\n", s1)
  fmt.Printf("Capacity: %v\n", cap(s1))
  fmt.Printf("Length: %v\n\n", len(s1))

  s1 = append(s1, 5)
  s1[0] = 1
  fmt.Printf("s1: %v\n", s1)
  fmt.Printf("Capacity: %v\n", cap(s1))
  fmt.Printf("Length: %v\n\n", len(s1))

  // Still point to old array.
  fmt.Printf("s2: %v\n", s2)
  fmt.Printf("Capacity: %v\n", cap(s2))
  fmt.Printf("Length: %v\n\n", len(s2))
}
