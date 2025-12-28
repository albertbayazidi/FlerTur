package main

import (
	"fmt"
	"time"
	"strconv"
)

func parseDateToUnix(dateStr string) (int64) {
	layout := "2006-01-02"

	unixTime, _ := time.Parse(layout, dateStr)
	unixTime = time.Date(unixTime.Year(), unixTime.Month(), unixTime.Day(), 0, 0, 0, 0, unixTime.Location())

	return unixTime.UnixMilli()
}

func metaDataify(dateStr string) string {
	metaDataLeft := "https://entur.no/reiseresultater?transportModes=rail%2Ctram%2Cbus%2Ccoach%2Cwater%2Ccar_ferry%2Cmetro%2Cflytog%2Cflybuss"
	metaDataRight :="&tripMode=oneway&walkSpeed=1.3&minimumTransferTime=120&timepickerMode=departAfter"

	return metaDataLeft + "&date=" + dateStr + metaDataRight 
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

func constructUrl(date string, start string, stop string) (string, error) {
	unixTime := parseDateToUnix(date)
	baseURL := metaDataify(strconv.FormatInt(unixTime, 10))
	
	startLatLongPair, err := checkStation(start, "start")
	if err != nil {
			return "", err
	}

	endLatLongPair, err := checkStation(stop, "stop")
	if err != nil {
			return "", err
	}

	url := baseURL + startLatLongPair + endLatLongPair
	return url, nil
}


