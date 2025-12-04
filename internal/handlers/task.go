package handlers

import (
	"encoding/json"
	"net/http"
	"task-api/internal/database"
	"task-api/internal/dto"
	"task-api/internal/models"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(models.Tasks)
	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	// idInt, _ := strconv.Atoi(id)

	// for _, task := range models.Tasks {
	// 	if task.ID == idInt {
	// 		json.NewEncoder(w).Encode(task)
	// 		return
	// 	}
	// }
	var task models.Task
	if result := database.DB.First(&task, id); result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)

	// http.Error(w, "Task not found", http.StatusNotFound)
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

	// task := models.Task{
	// 	ID:        models.NextID,
	// 	Title:     input.Title,
	// 	Done:      input.Done,
	// 	CreatedAt: time.Now(),
	// }
	// models.NextID++
	// models.Tasks = append(models.Tasks, task)
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(task)
	task := models.Task{
		Title: input.Title,
		Done:  input.Done,
	}
	database.DB.Create(&task)
	json.NewEncoder(w).Encode(task)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	// idInt, _ := strconv.Atoi(id)

	// for i, task := range models.Tasks {
	// 	if task.ID == idInt {
	// 		var updatedTask models.Task
	// 		json.NewDecoder(r.Body).Decode(&updatedTask)
	// 		updatedTask.ID = idInt
	// 		models.Tasks[i] = updatedTask
	// 		json.NewEncoder(w).Encode(updatedTask)
	// 		return
	// 	}
	// }
	// http.Error(w, "Task not found", http.StatusNotFound)

	var task models.Task
	if result := database.DB.First(&task, id); result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	var input dto.CreateTaskRequest
	json.NewDecoder(r.Body).Decode(&input)
	if errs := dto.ValidateCreateTask(input); errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": errs["Title"]})
		return
	}
	task.Title = input.Title
	task.Done = input.Done
	database.DB.Save(&task)
	json.NewEncoder(w).Encode(task)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// idInt, _ := strconv.Atoi(id)

	// for i, task := range models.Tasks {
	// 	if task.ID == idInt {
	// 		models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
	// 		w.WriteHeader(http.StatusNoContent)
	// 		return
	// 	}
	// }
	// http.Error(w, "Task not found", http.StatusNotFound)

	var task models.Task
	if result := database.DB.First(task, id); result.Error != nil {
		http.Error(w, "Tasks not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
