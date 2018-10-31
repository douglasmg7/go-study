package main

import (
  "fmt"
  "reflect"
)

func main() {
  ascii := 'a'
  fmt.Printf("%d %[1]c %[1]q (%[1]T)\n", ascii)
  fmt.Println(reflect.TypeOf(ascii))
  // Rune is alias for int32, not a defined type
  // An alias declaration
  // type T1 = T2
  // Type definition
  // type T1 T2
}
