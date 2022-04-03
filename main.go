package main

import (
	"github.com/kcm3394/choose-your-own-adventure/handler"
	"github.com/kcm3394/choose-your-own-adventure/model"
	"log"
	"net/http"
)

const (
	JSON_FILE = "gopher.json"
)

func main() {
	stories, err := model.JSONToStories(JSON_FILE)
	if err != nil {
		log.Fatalf("Failed to parse JSON. %v", err)
	}

	http.Handle("/", handler.NewStoryHandler(stories))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
