package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	item, err := insertTask(title)
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
	tmpl.ExecuteTemplate(w, "Item", map[string]any{"Item": item, "SwapOOB": true})
	tmpl.ExecuteTemplate(w, "TotalCount", map[string]any{"Count": count, "SwapOOB": true})
}

func handleToggleTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id into int %v", err)
		return
	}
	_, err = toggleTask(id)
	if err != nil {
		log.Printf("Error toggling task %v", err)
		return
	}
	completedCount, err := fetchCompletedCount()
	if err != nil {
		log.Printf("Error fetching completed count %v", err)
		return
	}
	tmpl.ExecuteTemplate(w, "CompletedCount", map[string]any{"Count": completedCount, "SwapOOB": true})
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id into int %v", err)
		return
	}

	err = deleteTask(r.Context(), id)
	if err != nil {
		log.Printf("Error deleting taks %v", err)
	}

	count, err := fetchCount()
	if err != nil {
		log.Printf("Error fetching count %v", err)
	}
	completedCount, err := fetchCompletedCount()
	if err != nil {
		log.Printf("Error fetching completed count %v", err)
	}

	tmpl.ExecuteTemplate(w, "TotalCount", map[string]any{"Count": count, "SwapOOB": true})
	tmpl.ExecuteTemplate(w, "CompletedCount", map[string]any{"Count": completedCount, "SwapOOB": true})
}

func handleEditTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id into int %v", err)
		return
	}

	task, err := fetchTask(id)
	if err != nil {
		log.Printf("Error parsing task with id %d %v", id, err)
		return
	}

	tmpl.ExecuteTemplate(w, "Item", map[string]any{"Item": task, "Editing": true})
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id into int %v", err)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		return
	}

	task, err := updateTask(id, title)
	if err != nil {
		log.Printf("Error fetching task with id %d %v", id, err)
		return
	}

	tmpl.ExecuteTemplate(w, "Item", map[string]any{"Item": task})
}
