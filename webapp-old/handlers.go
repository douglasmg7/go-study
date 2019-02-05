package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

// http://localhost:8080
func home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./template/home.page.tmpl",
		"./template/base.layout.tmpl",
		"./template/footer.partial.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Error", 500)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Error", 500)
	}
}

// http://localhost:8080/snippet?id=3
func showSnippet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// if err != nil || id < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }
	fmt.Fprintf(w, "Display a specific snippet with ID %s!\n", p.ByName("id"))
}

// curl -i -X POST http://localhost:8080/snippet/create
func createSnippet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet...\n"))
}
