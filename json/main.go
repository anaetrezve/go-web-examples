package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
	Age       int    `json: "age"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.Body)
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w, "%s %s is %d years old.", user.FirstName, user.LastName, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		peter := User{
			FirstName: "Jonh",
			LastName:  "Doe",
			Age:       25,
		}

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":5000", nil)
}
