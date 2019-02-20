package main

import "fmt"

func main() {
  m := map[string]int{}
  m["a"] = 1
  fmt.Printf("m: %v\n", m)
  fmt.Printf("have the key: %v\n", m["a"])
  fmt.Printf("No have the key: %v\n\n", m["b"])

  m2 := map[string]int{
    "b": 2,
    "c": 3,
  }
  fmt.Printf("m2: %v", m2)
}
