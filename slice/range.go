package main

import "fmt"

func main() {
  // Range.
  s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
  for i, v := range s1 {
    fmt.Printf("index: %v, value: %v\n", i, v)
  }
}
