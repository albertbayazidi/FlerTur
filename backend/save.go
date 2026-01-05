package main

import (
    "encoding/json"
		"time"
    "os"
		"backend/rod_utils"
)

type PageDataWrapper struct {
    StartStation       	string								`json:"startStation"`
    EndStation      		string    						`json:"endStation"`
    PageDataResults     []rod_utils.PageData  `json:"pageDataResults"`
    RetrievalTime  			time.Time 						`json:"retrievalTime"`
}

func SavePageData(filename string, records []rod_utils.PageData) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    
    encoder.SetIndent("", "  ") 
    
    return encoder.Encode(records)
}

func LoadPageData(filename string) ([]rod_utils.PageData, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var records []rod_utils.PageData
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&records)
    return records, err
}
