package main

import (
	"github.com/go-rod/rod"
	"time"
)

func scraper(url string) {
	page := rod.New().MustConnect().MustPage(url)
	page.MustWaitStable()
	page.MustElement("#main-content > div > section.travel-result__content__transit > button").MustClick() 
	//page.MustElement("#main-content > div > section.travel-result__content__transit > button").MustClick() 
	//page.MustElement("#main-content > div > section.travel-result__content__transit > button").MustClick() 

	time.Sleep(time.Hour)
}

