package main

import (
	"fmt"
	"time"
	"strconv"
)
var stationMap = map[string][]string{
		"Oslo S":                {"59.910357","10.753051"},
		"Trondheim S":           {"63.436279","10.399123"},
		"Stavanger stasjon":     {"58.966568","5.732616"},
		"Bergen stasjon":        {"60.390434","5.333511"},
		"Fredrikstad stasjon":   {"59.208805","10.950282"},
		"Kristiansand stasjon":  {"58.14559","7.988067"},
		"Asker stasjon":				 {"59.833128","10.434169"},
		"Arendal stasjon":			 {"58.465114","8.7693"},
    }

type Route struct {
	Start string
	End   string
}

var routes = []Route{
		{"Bergen stasjon", "Fredrikstad stasjon"},
		{"Bergen stasjon", "Kristiansand stasjon"},
		{"Bergen stasjon", "Asker stasjon"},

		{"Fredrikstad stasjon", "Kristiansand stasjon"},
		{"Fredrikstad Stasjon", "Asker stasjon"},

		{"Oslo S", "Bergen stasjon"},
		{"Oslo S", "Fredrikstad stasjon"},
		{"Oslo S", "Kristiansand stasjon"},
		{"Oslo S", "Stavanger stasjon"},
		{"Oslo S", "Trondheim S"},
		{"Oslo S", "Asker stasjon"},

		{"Stavanger stasjon", "Bergen stasjon"},
		{"Stavanger stasjon", "Fredrikstad stasjon"},
		{"Stavanger stasjon", "Kristiansand stasjon"},
		{"Stavanger stasjon", "Asker stasjon"},

		{"Trondheim S", "Fredrikstad stasjon"},
		{"Trondheim S", "Kristiansand stasjon"},
		{"Trondheim S", "Asker stasjon"},
		{"Trondheim S", "Arendal stasjon"},
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


