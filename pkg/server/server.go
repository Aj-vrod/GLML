package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GLML/pkg/movies"
	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", movies.MoviesHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", movies.MovieHandler).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
