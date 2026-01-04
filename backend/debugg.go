package main

import (
	"fmt"
)

func printPageData(data *pageData) {
	fmt.Println("===== pageData =====")
	fmt.Println("Duration:        ", data.Duration)
	fmt.Println("Start Time:      ", data.StartTime)
	fmt.Println("Price:           ", data.Price)
	fmt.Println("Number of Trains:", data.NumberOfTrains)
	fmt.Println("Url:             ", data.URL)
	fmt.Println("RetrievalTime   :", data.RetrievalTime)

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
