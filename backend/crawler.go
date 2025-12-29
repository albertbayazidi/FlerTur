package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func captureLiElement(page *rod.Page) (rod.Elements) {
	ulElements := page.MustElement("#main-content > div > section.travel-result__content__transit > div > div > ul")   
	liElements := ulElements.MustElements("li")

	return liElements
}

func pressButton(button *rod.Element, page *rod.Page) {
	i := 0
	liElemets := captureLiElement(page)
	nrElemets := len(liElemets)

	for nrElemets >= len(liElemets){
		liElemets = captureLiElement(page)
		time.Sleep(500*time.Millisecond)
		i++
		fmt.Println(i)
	} 
	button.MustClick()
}

func scraper(url string) {
	broswer := rod.New().MustConnect()
	page := broswer.MustPage(url)
	page.MustWaitStable()
	
	button := page.MustElement("#main-content > div > section.travel-result__content__transit > button")  
	
	button.MustClick()

	for {
		pressButton(button, page)

		if page.MustHas("#main-content > div > section.travel-result__content__transit > div > div:nth-child(2) > h3 > span") {

			break
		}
	}	
	fmt.Println("finished")

	time.Sleep(time.Hour)
}

