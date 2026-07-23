package main

import (
	"fmt"
	"sync"
	"time"

	"backend/rod_utils"
	"backend/types"

	"github.com/go-rod/rod"
)

var mu sync.Mutex
var maxDay = 6
var maxStaions = 8
var numThreads = 2
var restTime = 30

func mainProsses(browser *rod.Browser, startStation string, endStation string, currentDate string) types.PageDataWrapper {
	var pageDataResults []types.PageData

	currentDay := 0
	url, _ := constructUrl(currentDate, startStation, endStation)
	for currentDay <= maxDay {
		_ = browser.MustPage(url.Url)
		updateUrl(&url)
		currentDay++
	}

	currentDay = 0
	pageList, _ := browser.Pages()
	for _, currentPage := range pageList {
		currentPage.Activate()
		rod_utils.Crawler(currentPage)
		pageDataResults = append(pageDataResults, rod_utils.Scraper(currentPage))
		currentPage.MustClose()
	}

	wrapper := types.PageDataWrapper{
		StartStation:    startStation,
		EndStation:      endStation,
		RetrievalTime:   time.Now(),
		PageDataResults: pageDataResults,
	}
	return wrapper
}

func mainProssesSave() {
	db := ConnectDB()
	defer db.Close()

	var allResults []types.PageDataWrapper
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	currentDate := tomorrow.Format("2006-01-02")

	limit := maxStaions + 1
	if limit > len(routes) {
		limit = len(routes)
	}
	activeRoutes := routes[:limit]

	chunkSize := (len(activeRoutes) + numThreads - 1) / numThreads
	var wg sync.WaitGroup
	resultChan := make(chan types.PageDataWrapper, len(activeRoutes)*2)

	for i := 0; i < numThreads; i++ {
		startIdx := i * chunkSize
		if startIdx >= len(activeRoutes) {
			break
		}
		endIdx := startIdx + chunkSize
		if endIdx > len(activeRoutes) {
			endIdx = len(activeRoutes)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			browser := rod.New().MustConnect()
			defer browser.MustClose()

			for j := start; j < end; j++ {
				route := activeRoutes[j]
				fmt.Println("start", route.Start, "end", route.End)
				result1 := mainProsses(browser, route.Start, route.End, currentDate)
				resultChan <- result1

				time.Sleep(time.Second)

				result2 := mainProsses(browser, route.End, route.Start, currentDate)
				resultChan <- result2

				time.Sleep(time.Second)
			}
		}(startIdx, endIdx)
	}

	wg.Wait()
	close(resultChan)

	for res := range resultChan {
		allResults = append(allResults, res)
	}

	fmt.Println("Saving results to database")
	err := SaveToDB(db, allResults)
	if err != nil {
		fmt.Println("Fatal error saving to DB:", err)
	} else {
		fmt.Println("Successfully saved all routes!")
	}
}

func main() {
	interval := 30 * time.Minute

	for {
		start := time.Now()

		mainProssesSave()

		duration := time.Since(start)
		fmt.Printf("Process took: %v\n", duration)
		wait := interval - duration

		if wait > 0 {
			fmt.Printf("Waiting for %v before next run...\n", wait)
			time.Sleep(wait)
		} else {
			fmt.Println("Process took longer than 30 minutes. Restarting immediately.")
		}
	}
}
