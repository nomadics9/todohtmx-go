package main

import (
	"log"
	"net/http"
)

func handleGetTasks(w http.ResponseWriter, _ *http.Request) {
	items, err := fetchTasks()
	if err != nil {
		log.Printf("Error fetching Tasks %v", err)
		return
	}
	count, err := fetchCount()
	if err != nil {
		log.Printf("Error fetching Count %v", err)
	}
	completedCount, err := fetchCompletedCount()
	if err != nil {
		log.Printf("Error fetching CompletedCount %v", err)
	}
	data := Tasks{
		Items:          items,
		Count:          count,
		CompletedCount: completedCount,
	}
	tmpl.ExecuteTemplate(w, "Base", data)
}
