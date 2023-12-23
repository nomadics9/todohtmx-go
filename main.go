package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := openDB()
	if err != nil {
		log.Panic(err)
	}
	defer closeDB()

	err = setupDB()
	if err != nil {
		log.Panic(err)
	}

	err = parseTemplates()
	if err != nil {
		log.Panic(err)
	}

	server()
}

func server() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Get("/", handleGetTasks)

	http.ListenAndServe("localhost:3000", r)
}
