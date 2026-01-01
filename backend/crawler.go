package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

var newDayHeaderSelector = "#main-content > div > section.travel-result__content__transit > div > div:nth-child(2) > h3 > span" 
var buttonSelector = "#main-content > div > section.travel-result__content__transit > button" 

func captureLiElement(page *rod.Page) (rod.Elements) {
	ulElements := page.MustElement("#main-content > div > section.travel-result__content__transit > div > div > ul")   
	liElements := ulElements.MustElements("li")

	return liElements
}

func pressButton(button *rod.Element, page *rod.Page) bool {
	i := 0
	liElements := captureLiElement(page)
	nrElements := len(liElements)

	for nrElements >= len(liElements) {
		liElements = captureLiElement(page)
		time.Sleep(250 * time.Millisecond)

		if page.MustHas(newDayHeaderSelector) {
			return true
		}

		i++
		fmt.Println(i)
	}

	button.MustClick()
	return false
}

func scraper(page *rod.Page) {
	page.MustWaitStable()

	button := page.MustElement(buttonSelector)
	button.MustClick()

	for {
		shouldStop := pressButton(button, page)
		if shouldStop {
			break
		}
	}
	fmt.Println("finished page")
}

