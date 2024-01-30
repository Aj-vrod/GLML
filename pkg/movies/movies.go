package movies

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	id   int
	name string
}

func MoviesHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You reached Movies endpoint!\n")
}

func MovieHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]

	if !ok {
		fmt.Fprintf(w, "That's not a valid movie id :(")
		return
	}

	fmt.Fprintf(w, "You reached the %s movie endpoint!\n", id)
}
