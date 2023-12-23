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

func handleCreateTasks(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		tmpl.ExecuteTemplate(w, "Form", nil)
		return
	}
	_, err := insertTask(title)
	if err != nil {
		log.Printf("Error inserting Task %v", err)
		return
	}

	count, err := fetchCount()
	if err != nil {
		log.Printf("Error fetching Count %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	tmpl.ExecuteTemplate(w, "Form", nil)
	tmpl.ExecuteTemplate(w, "TotalCount", map[string]any{"Count": count, "SwapOOB": true})
}
