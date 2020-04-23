package main

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	bp, _ := filepath.Abs("./")
	fp := filepath.Join(bp, "forms")

	tmpl := template.Must(template.ParseFiles(filepath.Join(fp, "index.gohtml")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		formData := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		resp := struct {
			Success  bool
			FormData ContactDetails
		}{
			Success:  true,
			FormData: formData,
		}

		tmpl.Execute(w, resp)
	})

	http.ListenAndServe(":5000", nil)
}
