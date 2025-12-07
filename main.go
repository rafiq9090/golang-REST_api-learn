package main

import (
	"log"
	"net/http"
	"task-api/internal/database"
	"task-api/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	log.Println("Postgres connected")

	router := mux.NewRouter()

	// // Routes
	// router.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET")
	// router.HandleFunc("/api/tasks/{id}", handlers.GetTask).Methods("GET")
	// router.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")
	// router.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	// router.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	routes.Setup(router)
	log.Println("API চলছে → http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
