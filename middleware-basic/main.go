package main

import (
	"fmt"
	"log"
	"net/http"
)

func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		fn(w, r)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about")
}

func main() {
	http.HandleFunc("/", logger(indexPage))
	http.HandleFunc("/about", logger(aboutPage))

	fmt.Println("Server is running on port: 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
