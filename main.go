package main

import (
	"log"
	"net/http"
	"task-api/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("API চলছে → http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
