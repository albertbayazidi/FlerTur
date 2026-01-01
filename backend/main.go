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
		"Lillehammer stasjon":   {"61.114912","10.461479"},
		"Stavanger stasjon":     {"58.966568","5.732616"},
		"Bergen stasjon":        {"60.390434","5.333511"},
		"Lillestrøm stasjon":    {"59.952915","11.045364&"},
		"Fredrikstad stasjon":   {"59.208805","10.950282"},
		"Kristiansand stasjon":  {"58.14559","7.988067"},
		"Kongsvinger stasjon":   {"60.187497","12.004063"},
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
	maxDay := 7

	url, err := constructUrl(date,"Oslo S","Lillehammer stasjon")
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
	currentPage := pageList[currentDay]
	currentPage.Activate()

	for _, page := range pageList {
		page.Activate()
		scraper(page)
	}

	time.Sleep(time.Hour)



}
