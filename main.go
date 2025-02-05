package main

import (
	"net/http"
	"text/template"
	"voting-system/handlers"
)

func main() {
	http.HandleFunc("/vote", handlers.VoteHandler)
	http.HandleFunc("/results", handlers.ResultsHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/resultsPage", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/results.html"))
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}