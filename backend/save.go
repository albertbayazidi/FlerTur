package main

import (
	"os"

	"github.com/vmihailenco/msgpack/v5"
)

func SavePageData(filename string, records []pageData) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := msgpack.NewEncoder(file)
	return encoder.Encode(records)
}

func LoadPageData(filename string) ([]pageData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var records []pageData
	decoder := msgpack.NewDecoder(file)
	err = decoder.Decode(&records)
	return records, err
}
