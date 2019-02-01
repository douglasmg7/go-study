package main

import "fmt"

// Declare y as 0 because the var.
var y int

func main() {
  x := 7
  fmt.Printf("%T", x)
  fmt.Println(y)
}
