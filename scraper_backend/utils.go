package main

import (
	"fmt"
	"time"
	"strconv"
)
var stationMap = map[string][]string{
		"oslo S":                {"59.910357","10.753051"},
		"trondheim S":           {"63.436279","10.399123"},
		"stavanger stasjon":     {"58.966568","5.732616"},
		"bergen stasjon":        {"60.390434","5.333511"},
		"fredrikstad stasjon":   {"59.208805","10.950282"},
		"kristiansand stasjon":  {"58.14559","7.988067"},
		"asker stasjon":				 {"59.833128","10.434169"},
		"arendal stasjon":			 {"58.465114","8.7693"},
    }

type Route struct {
	Start string
	End   string
}

var routes = []Route{
		{"bergen stasjon", "fredrikstad stasjon"},
		{"bergen stasjon", "kristiansand stasjon"},
		{"bergen stasjon", "asker stasjon"},

		{"fredrikstad stasjon", "kristiansand stasjon"},
		{"fredrikstad stasjon", "asker stasjon"},

		{"oslo s", "bergen stasjon"},
		{"oslo s", "fredrikstad stasjon"},
		{"oslo s", "kristiansand stasjon"},
		{"oslo s", "stavanger stasjon"},
		{"oslo s", "trondheim s"},
		{"oslo s", "asker stasjon"},

		{"stavanger stasjon", "bergen stasjon"},
		{"stavanger stasjon", "fredrikstad stasjon"},
		{"stavanger stasjon", "kristiansand stasjon"},
		{"stavanger stasjon", "asker stasjon"},

		{"trondheim s", "fredrikstad stasjon"},
		{"trondheim s", "kristiansand stasjon"},
		{"trondheim s", "asker stasjon"},
		{"trondheim s", "arendal stasjon"},
	}

type urlAndMetaData struct {
	date time.Time
	stationLatLonPair string
	url string
}

var metaDataLeft = "https://entur.no/reiseresultater?transportModes=rail%2Ctram%2Cbus%2Ccoach%2Cwater%2Ccar_ferry%2Cmetro%2Cflytog%2Cflybuss"
var metaDataRight ="&tripMode=oneway&walkSpeed=1.3&minimumTransferTime=120&timepickerMode=departAfter"

func parseDateToTime(dateStr string) (time.Time) {
	layout := "2006-01-02"

	unixTime, _ := time.Parse(layout, dateStr)
	unixTime = time.Date(unixTime.Year(), unixTime.Month(), unixTime.Day(), 0, 0, 0, 0, unixTime.Location())

	return unixTime
}

func metaDataify(date time.Time) string {
	return metaDataLeft + "&date=" + strconv.FormatInt(date.UnixMilli(),10) + metaDataRight 
}

func checkStation(station string, sendState string) (string, error) {
	coords, ok := stationMap[station]
	if !ok {
			return "", fmt.Errorf("station not found: %s", station)
	}

	stationLatLonPair := "&" + sendState + "Label=" + station +
											 "&" + sendState + "Lat=" + coords[0] +
											 "&" + sendState + "Lon=" + coords[1]

	return stationLatLonPair, nil
}

func updateUrl(url *urlAndMetaData){
	url.date = url.date.AddDate(0,0,1)
	baseURL := metaDataify(url.date)
	url.url = baseURL + url.stationLatLonPair

}

func constructUrl(date string, start string, stop string) (urlAndMetaData, error) {
	timeTime := parseDateToTime(date)
	baseURL := metaDataify(timeTime)
	url := urlAndMetaData{date: timeTime}
	
	startLatLongPair, err := checkStation(start, "start")

	endLatLongPair, err := checkStation(stop, "stop")

	url.stationLatLonPair= startLatLongPair + endLatLongPair 
	url.url = baseURL + startLatLongPair + endLatLongPair

	return url, err
}


