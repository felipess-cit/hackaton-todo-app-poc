package main

import (
	"html/template"
	"log"
	"net/http"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks = []Task{
	{ID: 1, Description: "Buy milk", Completed: false},
	{ID: 2, Description: "Read a book", Completed: true},
	{ID: 3, Description: "Write Go app", Completed: false},
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Server starting on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}