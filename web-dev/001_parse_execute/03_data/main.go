package main

import (
  "log"
  "os"
  "text/template"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("index.gohtml", "index2.gohtml"))
}

func main() {
  err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 777)
  if err != nil {
    log.Fatalln(err)
  }
  err = tpl.ExecuteTemplate(os.Stdout, "index2.gohtml", "Ula ula")
  if err != nil {
    log.Fatalln(err)
  }
}
