package main

import (
	"fmt"
	"backend/rod_utils"

)

func PrintPageDataWrapper(wrapper PageDataWrapper) {
    fmt.Println("=====SEARCH RESULTS=====")
    fmt.Printf(" Route:    %s  -->  %s\n", wrapper.StartStation, wrapper.EndStation)
    fmt.Printf(" Fetched:  %s\n", wrapper.RetrievalTime.Format("2006-01-02 15:04:05"))
    fmt.Println("========================")

    if len(wrapper.PageDataResults) == 0 {
        fmt.Println(" [!] No train data found for this route.")
    }

    for i, record := range wrapper.PageDataResults {
        fmt.Printf("\n--- Day %d ---\n", i+1)
        printPageData(&record)
    }
    
    fmt.Println("========================")
}

func printPageData(data *rod_utils.PageData) {
	fmt.Println("===== pageData =====")
	fmt.Println("Duration:        ", data.Duration)
	fmt.Println("Start Time:      ", data.StartTime)
	fmt.Println("Price:           ", data.Price)
	fmt.Println("Number of Trains:", data.NumberOfTrains)
	fmt.Println("Url:             ", data.URL)

	fmt.Println("Train IDs:")
	for i, id := range data.TrainIds {
		fmt.Printf("  %d: %s\n", i+1, id)
	}

	fmt.Println("====================")
}

func stopCodeUntilPres(){
	var cont string
	fmt.Print("Press any button to contionue: ")
	fmt.Scanln(cont)
	fmt.Println()
}
