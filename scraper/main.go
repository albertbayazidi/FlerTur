package main

import (
	"time"
	"sync"
	"github.com/go-rod/rod"
	"fmt"
	"backend/rod_utils"
)

var mu sync.Mutex 
var maxDay = 6 // keep at low for testing  //Final number should be 13
var maxStaions = 8 // keep at low for testing  // Final number should be 8

func mainProsses(startStation string, endStation string, currentDate string) PageDataWrapper{
	browser := rod.New().MustConnect()
	defer browser.MustClose() 	
	
	var pageDataResults []rod_utils.PageData 

	currentDay := 0
	url, _ := constructUrl(currentDate, startStation, endStation)
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
	return wrapper
}

func mainProssesSave(){
	db := ConnectDB()
  defer db.Close()

	var allResults []PageDataWrapper
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	currentDate := tomorrow.Format("2006-01-02")

	for index, route := range routes {
		result1 := mainProsses(route.Start, route.End, currentDate)
		allResults = append(allResults, result1)

		time.Sleep(time.Second) 

		result2 := mainProsses(route.End, route.Start, currentDate)
		allResults = append(allResults, result2)
		
		time.Sleep(time.Second)

		if index == maxStaions {
			break
		}
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
