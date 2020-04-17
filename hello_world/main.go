package main

// Here we are importing 'fmt' and http library package
import (
	"fmt"
	"net/http"
)

// Main function will run when we start the app
func main() {

	// we are creating request handler which will handle all requests
	// coming to server at '/' url
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Here we are writing hello world message to response writer
		// To send back to client
		_, err := fmt.Fprintf(w, "Hello World! From Golang Server.")

		// Checking if there is any error
		if err != nil {
			// If has any error then return Internal Server Error to client
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	// Our server will start and listen on localhost:5000
	// Any request come on this url our server will accecpt the request
	http.ListenAndServe("localhost:5000", nil)
}
