package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler function used as a callback
func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	// setting up the routes & handlers
	http.HandleFunc("/", helloFunc)
	// starting the server & making it listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
