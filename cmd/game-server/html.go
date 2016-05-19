package main

import (
	"html/template"
	"net/http"

	"github.com/bign8/games"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, games.List())
}

var rootTemplate = template.Must(template.ParseFiles("www/index.html"))
