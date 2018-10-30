package main

import (
  "fmt"
)

var buffer [12]byte

// type path []byte

func main() {
  // Init buffer.
  for i := range buffer {
    buffer[i] = byte(i)
  }

  fmt.Println("buffer:", buffer)
  slice := buffer[5:12]
  slice[0] = 30
  fmt.Println("slice :", slice)
  fmt.Println("buffer:", buffer)
}
