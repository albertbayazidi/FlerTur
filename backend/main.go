package main

import (
	"sync"
	"fmt"
)

var stationMap = map[string][]string{
		"Oslo%20S":                {"59.910357","10.753051"},
		"Trondheim%20S":           {"63.436279","10.399123"},
		"Lillehammer%20stasjon":   {"61.114912","10.461479"},
		"Stavanger%20stasjon":     {"58.966568","5.732616"},
		"Bergen%20stasjon":        {"60.390434","5.333511"},
		"Lillestrøm%20stasjon":    {"59.952915","11.045364&"},
		"Fredrikstad%20stasjon":   {"59.208805","10.950282"},
		"Kristiansand%20stasjon":  {"58.14559","7.988067"},
		"Kongsvinger%20stasjon":   {"60.187497","12.004063"},
    }

type ScrapeResult struct {
    URL       string `json:"url"`
    Start     string `json:"start"`
    End       string `json:"end"`
    Date      string `json:"date"`
    Price     string `json:"price"`
    Time      string `json:"time"`
}

var store []ScrapeResult
var mu sync.Mutex 

func main() {
	date := "2026-01-06"

	url, err := constructUrl(date,"Oslo%20S","Lillehammer%20stasjon")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	scraper(url)
}
