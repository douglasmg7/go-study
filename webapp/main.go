package main

import (
	"log"
	"net/http"
)

func main() {
	// redirection
	rh := http.RedirectHandler("https://www.zunka.com.br", 307)

	mux := http.NewServeMux()
	mux.Handle("/zunka", rh)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal("ListenAndServe", err)
	// if err != nil {
	// 	log.Fatal("ListenAndServe", err)
	// }
}

// go run *
