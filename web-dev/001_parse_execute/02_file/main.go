package main

import (
  "log"
  "os"
  "text/template"
)

func main() {
  // Parse template file.
  tpl, err := template.ParseFiles("index.gohtml")
  if err != nil {
    log.Fatalln(err)
  }

  // Create a new file.
  nf, err := os.Create("index.html")
  if err != nil {
    log.Println("error creating file", err)
  }
  defer nf.Close()

  // Create file from template.
  err = tpl.Execute(nf, nil)
  if err != nil {
    log.Fatalln(err)
  }
}
