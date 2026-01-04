package main

import (
	"fmt"
)

func printPageData(data *pageData) {
	fmt.Println("===== pageData =====")
	fmt.Println("Duration:        ", data.duration)
	fmt.Println("Start Time:      ", data.startTime)
	fmt.Println("Price:           ", data.price)
	fmt.Println("Number of Trains:", data.numberOfTrains)
	fmt.Println("Url:             ", data.url)
	fmt.Println("RetrievalTime   :", data.retrievalTime)

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
