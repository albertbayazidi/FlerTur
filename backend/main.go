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

var pageDataResults []pageData 
var mu sync.Mutex 

func main() {
	date := "2026-01-06" // should be an input read from browser
	browser := rod.New().MustConnect()
	
	currentDay := 0
	maxDay := 0 // keep at 0 for testing

	url, err := constructUrl(date,"Trondheim S","Oslo S")
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
		pageDataResults = append(pageDataResults, scraper(currentPage))
	}

	SavePageData("data/pageData.msgpack", pageDataResults)
	loaded, _ := LoadPageData("data/pageData.msgpack")

	for _, pageDataResult := range loaded {
		printPageData(&pageDataResult)
	}

	time.Sleep(time.Hour)
}
