package main

import (
	_ "github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	// redirection
	rh := http.RedirectHandler("https://www.zunka.com.br", 307)

	// mux_ := httprouter.New()
	mux := http.NewServeMux()
	mux.Handle("/zunka", rh)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal("ListenAndServe", err)
	// if err != nil {
	// 	log.Fatal("ListenAndServe", err)
	// }
}

// go run *
