package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	// open the file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// file close
	// function return.
	defer file.Close()

	// decode the file info a slice of pointers to feed values
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	//
	return feeds, err
}
