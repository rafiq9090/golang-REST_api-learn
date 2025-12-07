package routes

import (
	"task-api/internal/handlers"

	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")

}
