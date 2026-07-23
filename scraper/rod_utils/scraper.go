package rod_utils

import (
	"regexp"
	"strconv"
	"strings"

	"backend/types"

	"github.com/go-rod/rod"
)

const maxNumberOfTrains = 10

var durationSelector = "div.transit-result-item__content > h2 > span.transit-result-item__header__duration"

var startTimeSelector = "div.transit-result-item__content > div > div.legs-list > ul > li:nth-child(2) > div.legs-list__leg__time > time"

var priceSelector = "div.transit-result-item__footer > span:nth-child(1)"

var trainIdContainerSelector = "div.transit-result-item__content > div > div.legs-list > ul > li.legs-list__leg"

var idSelector = "div.legs-list__leg__details > div.travel-tag > span.travel-tag__label"

func captureTrainId(travelSuggestion *rod.Element, data *types.PageData) {
	count := 0
	var trainIdArray [maxNumberOfTrains]string
	trainIdContainer := travelSuggestion.MustElements(trainIdContainerSelector)

	for _, trainIdLi := range trainIdContainer {
		trainIdElements := trainIdLi.MustElements(idSelector)

		for _, trainIdElement := range trainIdElements {
			trainId, _ := trainIdElement.Text()

			if trainId != "" {
				trainId = strings.Split(trainId, " ")[0]
				trainIdArray[count] = trainId
				count++
			}
		}
	}
	data.NumberOfTrains = count
	data.TrainIds = trainIdArray[:count]
}

func capturePrice(priceString string) int {
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

func captureUrl(page *rod.Page, travelSuggestions rod.Elements, currentCheapestTicketIndex int) string {
	travelSuggestions[currentCheapestTicketIndex].MustClick()
	return page.MustInfo().URL
}

func safeText(el *rod.Element, selector string) string {
	child, err := el.Element(selector)
	if err != nil || child == nil {
		return ""
	}

	text, err := child.Text()
	if err != nil {
		return ""
	}
	return text
}

func captureData(travelSuggestion *rod.Element) types.PageData {
	duration := safeText(travelSuggestion, durationSelector)
	startTime := safeText(travelSuggestion, startTimeSelector)
	priceString := safeText(travelSuggestion, priceSelector)

	data := types.PageData{
		Duration:  duration,
		StartTime: startTime,
		Price:     capturePrice(priceString),
	}

	captureTrainId(travelSuggestion, &data)
	return data
}

func Scraper(page *rod.Page) types.PageData {
	travelSuggestions := captureLiElements(page)
	currentCheapestTicketIndex := 0
	currentCheapestTicket := 9999999
	var cheapestData types.PageData

	for index, travelSuggestion := range travelSuggestions {
		data := captureData(travelSuggestion)

		if data.Price < currentCheapestTicket {
			currentCheapestTicket = data.Price
			currentCheapestTicketIndex = index
			cheapestData = data
		}
	}

	if len(travelSuggestions) > 0 {
		cheapestData.URL = captureUrl(page, travelSuggestions, currentCheapestTicketIndex)
	}

	return cheapestData
}
