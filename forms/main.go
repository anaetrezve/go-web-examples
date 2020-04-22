package main

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func main() {
	bp, _ := filepath.Abs("./")
	fp := filepath.Join(bp, "forms")

	tmpl := template.Must(template.ParseFiles(filepath.Join(fp, "index.gohtml")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":5000", nil)
}
