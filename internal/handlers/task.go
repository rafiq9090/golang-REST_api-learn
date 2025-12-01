package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-api/internal/dto"
	"task-api/internal/models"
	"time"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	for _, task := range models.Tasks {
		if task.ID == idInt {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var task models.Task
	// json.NewDecoder(r.Body).Decode(&task)
	// task.ID = models.NextID
	// models.NextID++
	// models.Tasks = append(models.Tasks, task)
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(task)
	var input dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	if errs := dto.ValidateCreateTask(input); errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Validation failed", "details": errs["Title"]})
		return
	}

	task := models.Task{
		ID:        models.NextID,
		Title:     input.Title,
		Done:      input.Done,
		CreatedAt: time.Now(),
	}
	models.NextID++
	models.Tasks = append(models.Tasks, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	for i, task := range models.Tasks {
		if task.ID == idInt {
			var updatedTask models.Task
			json.NewDecoder(r.Body).Decode(&updatedTask)
			updatedTask.ID = idInt
			models.Tasks[i] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	for i, task := range models.Tasks {
		if task.ID == idInt {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
