package main

import (
	_ "database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

const port = ":8080"

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := httprouter.New()
	router.GET("/favicon.ico", faviconHandler)
	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/contact", contact)
	router.GET("/apply", apply)
	router.POST("/apply", applyProcess)
	router.GET("/user/:name", user)
	router.GET("/blog/:category/:article", blogRead)
	router.POST("/blog/:category/:article", blogWrite)

	router.GET("/err", errExample)

	router.ServeFiles("/static/*filepath", http.Dir("./static/"))

	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func errExample(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.Error(w, "Some thing wrong", 404)
	return
}

func faviconHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(w, req, "./static/img/favicon.ico")
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}

func blogWrite(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
}

func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

// go run *
