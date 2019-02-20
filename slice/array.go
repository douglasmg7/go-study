package main

import (
  "fmt"
)

// Array.
// Documentation said, use slice instead.
// Len is part of type.
var a [5]int

func main() {
  a[1] = 2
  fmt.Println(a)
  fmt.Println(len(a))
}
