package main

import (
	"log"
	"net/http"
	"task-api/internal/config"
	"task-api/internal/database"
	"task-api/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.Load()
	database.Connect()
	log.Printf("%s started on port %s", config.App.AppName, config.App.Port)

	router := mux.NewRouter()

	// // Routes
	// router.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET")
	// router.HandleFunc("/api/tasks/{id}", handlers.GetTask).Methods("GET")
	// router.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")
	// router.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	// router.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	routes.Setup(router)
	log.Println("API চলছে → http://localhost:" + config.App.Port)
	log.Fatal(http.ListenAndServe(":"+config.App.Port, router))
}
