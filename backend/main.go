package main

import (
	"time"
	"sync"
	"fmt"
	"github.com/go-rod/rod"
)

var stationMap = map[string][]string{
		"Oslo S":                {"59.910357","10.753051"},
		"Trondheim S":           {"63.436279","10.399123"},
		"Stavanger stasjon":     {"58.966568","5.732616"},
		"Bergen stasjon":        {"60.390434","5.333511"},
		"Fredrikstad stasjon":   {"59.208805","10.950282"},
		"Kristiansand stasjon":  {"58.14559","7.988067"},
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
	date := "2026-01-06" // should be an input read from browser
	browser := rod.New().MustConnect()
	
	currentDay := 0
	maxDay := 0 // keep at 0 for testing

	url, err := constructUrl(date,"Trondheim S","Kongsvinger stasjon")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	
	browser = rod.New().MustConnect()
	for currentDay <=  maxDay {
		_ = browser.MustPage(url.url)
		updateUrl(&url)
		currentDay++
	}

	currentDay = 0
	pageList, _ := browser.Pages()
	for _, currentPage := range pageList {
		currentPage.Activate()
		crawler(currentPage)
		scraper(currentPage)
	}
	time.Sleep(time.Hour)

}
