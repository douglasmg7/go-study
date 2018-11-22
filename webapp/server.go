package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Form", r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Println()
	fmt.Fprintf(w, "Hello World")
}

func aResp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from aResp."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHelloName)
	mux.HandleFunc("/a", aResp)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal("ListenAndServe", err)
	// if err != nil {
	// 	log.Fatal("ListenAndServe", err)
	// }
}
