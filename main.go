package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
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

var idCounter = 3
var mu sync.Mutex

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addTaskHandler)
    http.Handle("/assets", http.FileServer(http.Dir("./images")))
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

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	description := r.FormValue("description")
	mu.Lock()
	idCounter++
	tasks = append(tasks, Task{ID: idCounter, Description: description, Completed: false})
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}