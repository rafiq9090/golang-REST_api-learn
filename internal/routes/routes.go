package routes

import (
	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	TaskRoutes(api)
	ProductRoutes(api)
}
