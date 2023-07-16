package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

type Todo struct { // Fix 1: Use "struct" instead of "struck"
	Item string
	Done bool
}

type PageData struct { // Fix 1: Use "struct" instead of "struck"
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn GO", Done: false},
			{Item: "Like this video", Done: false}, // Fix 2: Add a comma here
		}, // Fix 2: Add a comma after the third item
	}

	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/indexgo.html"))

	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs)) // Use mux.Handle here
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":9091", mux))
}
