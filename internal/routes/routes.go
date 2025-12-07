package routes

import (
	"task-api/internal/middleware"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.CORSMiddleware)
	api.Use(middleware.RateLimit)
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.JWTAuthMiddleware)
	TaskRoutes(protected)
	ProductRoutes(api)
	AuthRoutes(api)
}
