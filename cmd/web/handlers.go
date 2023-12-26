package main

import (
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	parseFile, err := template.ParseFiles("../../ui/html/index.html")
	if err != nil {
		return
	}
	err = parseFile.Execute(w, nil)
	if err != nil {
		return
	}
}
