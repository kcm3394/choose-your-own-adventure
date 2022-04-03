package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Stories map[string]*Chapter

type Chapter struct {
	Title   string     `json:"title"`
	Story   []string   `json:"story"`
	Options []*Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JSONToStories(file string) (Stories, error) {
	var stories Stories
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &stories)
	if err != nil {
		return nil, err
	}

	return stories, nil
}
