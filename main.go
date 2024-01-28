package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func startServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, GLML\n")
}

func main() {
	fmt.Println("Starting server...")
	fmt.Println("Listening on port :8080")

	r := mux.NewRouter()
	r.HandleFunc("/", startServer).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
