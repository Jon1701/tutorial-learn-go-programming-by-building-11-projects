package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new Router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	// Serve on port 9010
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
