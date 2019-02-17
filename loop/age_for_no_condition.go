package main

import "fmt"

func main() {
  bd := 1976
  for {
    if bd > 2019 {
      break
    }
    fmt.Println(bd)
    bd++
  }
}
