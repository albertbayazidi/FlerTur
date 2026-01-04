package main

import (
	"strings"
	"regexp"
	"strconv"
	"time"

	"github.com/go-rod/rod"
)

type pageData struct {
	duration string
	startTime string
	price int
	numberOfTrains int
	TrainIds []string
	url string
	retrievalTime time.Time
}

const maxNumberOfTrains = 10

var durationSelector = "div.transit-result-item__content > h2 > span.transit-result-item__header__duration" 

var startTimeSelector = "div.transit-result-item__content > div > div.legs-list > ul > li:nth-child(2) > div.legs-list__leg__time > time"

var priceSelector = "div.transit-result-item__footer > span:nth-child(1)"

var trainIdContainerSelector = "div.transit-result-item__content > div > div.legs-list > ul > li.legs-list__leg"

var idSelector = "div.legs-list__leg__details > div.travel-tag > span.travel-tag__label"

func captureTrainId(travelSuggestion *rod.Element, data *pageData) {
	count := 0
	var trainIdArray [maxNumberOfTrains]string
	trainIdContainer := travelSuggestion.MustElements(trainIdContainerSelector)

	for _, trainIdLi := range trainIdContainer{
		trainIdElements := trainIdLi.MustElements(idSelector)

		for _, trainIdElement := range trainIdElements{
			trainId, _ := trainIdElement.Text()

				if trainId != "" {
					trainId = strings.Split(trainId, " ")[0]
					trainIdArray[count] = trainId
					count++
			}
		}
	}	
	data.numberOfTrains = count
	data.TrainIds = trainIdArray[:count]
}

func capturePrice(priceString string) int{
	re := regexp.MustCompile(`(\d+)\s*kr`)
	match := re.FindStringSubmatch(priceString)

	if len(match) < 2 {
		return 9999999
	}
	price, err := strconv.Atoi(match[1])
	if err != nil {
		return 9999999
	}
	return price
}

func captureData(travelSuggestion *rod.Element) pageData {
	duration, _ := travelSuggestion.MustElement(durationSelector).Text()
	startTime ,_ := travelSuggestion.MustElement(startTimeSelector).Text()
	priceString, _ := travelSuggestion.MustElement(priceSelector).Text()
	
	data := pageData{duration: duration}
	data.startTime = startTime
	data.price = capturePrice(priceString)
	captureTrainId(travelSuggestion, &data)
	return data
}

func captureUrl(page *rod.Page, travelSuggestions rod.Elements, currentCheapestTicketIndex int) string {
	travelSuggestions[currentCheapestTicketIndex].MustClick()
	return page.MustInfo().URL
}

func scraper(page *rod.Page){
	travelSuggestions := captureLiElements(page)
	currentCheapestTicketIndex := 0
	currentCheapestTicket := 9999999

	for index, travelSuggestion := range travelSuggestions{
		data := captureData(travelSuggestion)

		if data.price < currentCheapestTicket{
			currentCheapestTicket = data.price 
			currentCheapestTicketIndex = index
		} 
	}
	
	data := captureData(travelSuggestions[currentCheapestTicketIndex])
	data.url = captureUrl(page,travelSuggestions,currentCheapestTicketIndex)
	data.retrievalTime = time.Now()
	printPageData(&data)

}

