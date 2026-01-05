package main

import (
	"time"
	"sync"
	"fmt"
	"github.com/go-rod/rod"
	"os"
	"backend/rod_utils"
)

var stationMap = map[string][]string{
		"Oslo S":                {"59.910357","10.753051"},
		"Trondheim S":           {"63.436279","10.399123"},
		"Stavanger stasjon":     {"58.966568","5.732616"},
		"Bergen stasjon":        {"60.390434","5.333511"},
		"Fredrikstad stasjon":   {"59.208805","10.950282"},
		"Kristiansand stasjon":  {"58.14559","7.988067"},
    }

var pageDataResults []rod_utils.PageData 
var mu sync.Mutex 
var maxDay = 1  // keep at low for testing 

func main() {
	startStation := os.Args[1]
	endStation := os.Args[2] 

	now := time.Now()
	currentDate := now.Format("2006-01-02")

	browser := rod.New().MustConnect()
	currentDay := 0

	url, err := constructUrl(currentDate, startStation, endStation)
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
		rod_utils.Crawler(currentPage)
		pageDataResults = append(pageDataResults, rod_utils.Scraper(currentPage))
	}

	wrapper := PageDataWrapper{StartStation: startStation,
																						EndStation: endStation,
																						RetrievalTime: time.Now(),
																						PageDataResults: pageDataResults}
	PrintPageDataWrapper(wrapper)

	/*
	SavePageData("../public/pageData.json", wrappedPageDataResults)
	loaded, _ := LoadPageData("../public/pageData.json")

	for _, pageDataResult := range loaded {
		printPageData(&pageDataResult)
	}
	*/

	time.Sleep(time.Hour)
}
