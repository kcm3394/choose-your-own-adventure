package handler

import (
	"github.com/kcm3394/choose-your-own-adventure/model"
	"html/template"
	"log"
	"net/http"
)

var storyTmpl = template.Must(template.ParseFiles("static/story-layout.html"))
var errorTmpl = template.Must(template.ParseFiles("static/error-layout.html"))

func NewStoryHandler(s model.Stories) http.Handler {
	return storyHandler{s}
}

type storyHandler struct {
	s model.Stories
}

func (sh storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if path == "" {
		story := sh.s["intro"]
		serveTemplate(w, storyTmpl, &story)
		return
	}

	story, ok := sh.s[path]
	if !ok {
		log.Printf("Unknown path: %s", path)
		w.WriteHeader(http.StatusNotFound)
		errorResp := &model.Error{
			StatusCode: http.StatusNotFound,
			Message:    "Not found: Unknown path",
		}
		serveTemplate(w, errorTmpl, errorResp)
		return
	}

	serveTemplate(w, storyTmpl, &story)
}

func serveTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Failed to write response. %v", err)
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
	}
}
