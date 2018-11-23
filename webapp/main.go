package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.GET("/", home)
	router.GET("/snippet/:id", showSnippet)
	router.POST("/snippet/create", createSnippet)

	router.ServeFiles("/static/*filepath", http.Dir("./static/"))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal("ListenAndServe", err)
}

// go run *
