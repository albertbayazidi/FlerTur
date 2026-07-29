package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"backend/rod_utils"
	"backend/types"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var maxDay int
var maxStations int
var numThreads int
var restTime int

func getEnvAsInt(name string, defaultValue int) int {
	valStr := os.Getenv(name)
	if valStr != "" {
		if val, err := strconv.Atoi(valStr); err == nil {
			return val
		}
	}
	return defaultValue
}

func init() {
	maxDay = getEnvAsInt("MAX_DAY", 6)
	maxStations = getEnvAsInt("MAX_STATIONS", 8)
	numThreads = getEnvAsInt("NUM_THREADS", 2)
	restTime = getEnvAsInt("REST_TIME", 30)
}

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

	limit := min(maxStations+1, len(routes))
	activeRoutes := routes[:limit]

	chunkSize := (len(activeRoutes) + numThreads - 1) / numThreads
	var wg sync.WaitGroup
	resultChan := make(chan types.PageDataWrapper, len(activeRoutes)*2)

	for i := range numThreads {
		startIdx := i * chunkSize
		if startIdx >= len(activeRoutes) {
			break
		}
		endIdx := min(startIdx+chunkSize, len(activeRoutes))

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			binPath := os.Getenv("ROD_BROWSER_BIN")
			l := launcher.New()
			if binPath != "" {
				l = l.Bin(binPath)
			}
			u := l.Set("no-sandbox").MustLaunch()

			browser := rod.New().ControlURL(u).MustConnect()
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
	interval := time.Duration(restTime) * time.Minute

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
			fmt.Println("Process took longer than interval. Restarting immediately.")
		}
	}
}
