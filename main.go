package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// To be implemented.
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// To be implemented.
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// To be implemented.
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// To be implemented.
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
}
