package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl, _ = template.ParseGlob("Template/*.html")
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "hello there")
	tpl.ExecuteTemplate(w, "index.html", nil)
	// tpl.ParseFiles("index.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "hello there")
	tpl.ExecuteTemplate(w, "about.html", nil)
	// tpl.ParseFiles("index.html")
}
