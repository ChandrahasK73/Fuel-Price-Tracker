package scraper

import (
    "net/http"
    "io/ioutil"
    "log"
)

func GetFuelPrices(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error fetching fuel prices: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %v", err)
    }
    return string(body) // This can be further processed
}