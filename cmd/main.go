package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/fuelprices", GetFuelPrices).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", r))
}

func GetFuelPrices(w http.ResponseWriter, r *http.Request) {
    // Implementation here
}
