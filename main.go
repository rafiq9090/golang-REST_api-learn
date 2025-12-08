// @title           Task Manager API
// @version         1.0
// @description     A simple task management API with JWT authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name   MIT
// @license.url    https://opensource.org/licenses/MIT

// @host      localhost:8081
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"log"
	"net/http"
	"task-api/internal/config"
	"task-api/internal/database"
	"task-api/internal/routes"

	_ "task-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.Load()
	database.Connect()
	log.Printf("%s started on port %s", config.App.AppName, config.App.Port)

	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
	))

	// // Routes
	// router.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET")
	// router.HandleFunc("/api/tasks/{id}", handlers.GetTask).Methods("GET")
	// router.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")
	// router.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	// router.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	routes.Setup(router)
	log.Println("Swagger UI: http://localhost:8081/swagger/index.html")
	log.Println("API running → http://localhost:" + config.App.Port)
	log.Fatal(http.ListenAndServe(":"+config.App.Port, router))
}
