package main

import (
	"log"
	"net/http"

	"github.com/angelospanag/council_tax_scraper/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/addresses/{postcode}", handlers.GetResults).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
