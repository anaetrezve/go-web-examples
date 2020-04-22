package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Todo structure of a single todo
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData is will be passed to template
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tp, _ := filepath.Abs("./templates_files")
	tmpl := template.Must(template.ParseFiles(tp + "/" + "index.tmpl"))

	todos := []Todo{
		{Title: "Todo One", Done: true},
		{Title: "Todo Two", Done: false},
		{Title: "Todo Three", Done: false},
		{Title: "Todo Four", Done: true},
	}

	data := TodoPageData{
		PageTitle: "Todo Page",
		Todos:     todos,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":5000", nil)
}
