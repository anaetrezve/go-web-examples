package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// PageData is for template data
type PageData struct {
	PageTitle string
}

func main() {
	path, _ := filepath.Abs("./")
	tp := filepath.Join(path, "static_files")
	fs := http.FileServer(http.Dir(filepath.Join(tp, "assets", "/")))

	indexPage := filepath.Join(tp, "templates", "index.tmpl")

	tmpl := template.Must(template.ParseFiles(indexPage))

	data := PageData{
		PageTitle: "Static File Serving with Golang",
	}

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":5000", nil)
}
