package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv( "PORT")

	http.HandleFunc("/hello", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "World")
	})

	http.ListenAndServe(":"+port, nil)
}