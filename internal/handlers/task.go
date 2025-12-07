package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-api/internal/dto"
	"task-api/internal/service"
	"task-api/internal/utils"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// // json.NewEncoder(w).Encode(models.Tasks)
	// var tasks []models.Task
	// database.DB.Find(&tasks)
	// json.NewEncoder(w).Encode(tasks)
	tasks, err := service.Task.GetAll()
	if err != nil {
		utils.JSONError(w, "Failed to fetch tasks", http.StatusInternalServerError, nil)
		// http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	utils.JSONSuccess(w, tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	// idInt, _ := strconv.Atoi(id)

	// for _, task := range models.Tasks {
	// 	if task.ID == idInt {
	// 		json.NewEncoder(w).Encode(task)
	// 		return
	// 	}
	// }
	// var task models.Task
	// if result := database.DB.First(&task, id); result.Error != nil {
	// 	http.Error(w, "Task not found", http.StatusNotFound)
	// 	return
	// }
	// json.NewEncoder(w).Encode(task)
	idInt, _ := strconv.ParseUint(id, 10, 64)
	task, err := service.Task.GetByID(uint(idInt))
	if err != nil {
		utils.JSONError(w, "Task not found", http.StatusNotFound, nil)
		return
	}
	utils.JSONSuccess(w, task)
	// http.Error(w, "Task not found", http.StatusNotFound)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	// var task models.Task
	// json.NewDecoder(r.Body).Decode(&task)
	// task.ID = models.NextID
	// models.NextID++
	// models.Tasks = append(models.Tasks, task)
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(task)
	var input dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "Invalid request body", http.StatusBadRequest, nil)
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	if errs := dto.ValidateCreateTask(input); errs != nil {
		utils.JSONError(w, "Validation failed", http.StatusBadRequest, map[string]string{"error": "Validation failed", "details": errs["Title"]})
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(map[string]string{"error": "Validation failed", "details": errs["Title"]})
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
	task, errs := service.Task.Create(input.Title, input.Done)
	if errs != nil {
		utils.JSONError(w, "Failed to create task", http.StatusInternalServerError, nil)
		return
	}
	utils.JSONSuccess(w, task)
	// task := models.Task{
	// 	Title: input.Title,
	// 	Done:  input.Done,
	// }
	// database.DB.Create(&task)
	// json.NewEncoder(w).Encode(task)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.ParseUint(id, 10, 64)
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

	// var task models.Task
	// if result := database.DB.First(&task, id); result.Error != nil {
	// 	http.Error(w, "Task not found", http.StatusNotFound)
	// 	return
	// }
	var input dto.CreateTaskRequest
	// json.NewDecoder(r.Body).Decode(&input)
	// if errs := dto.ValidateCreateTask(input); errs != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(map[string]string{"error": errs["Title"]})
	// 	return
	// }
	// task.Title = input.Title
	// task.Done = input.Done
	// database.DB.Save(&task)
	// json.NewEncoder(w).Encode(task)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "Invalid request body", http.StatusBadRequest, nil)
		return
	}
	if errs := dto.ValidateCreateTask(input); errs != nil {
		utils.JSONError(w, "Validation failed", http.StatusBadRequest, map[string]string{"error": "Validation failed", "details": errs["Title"]})
		return
	}
	task, err := service.Task.GetByID(uint(idInt))
	if err != nil {
		utils.JSONError(w, "Failed to update task", http.StatusInternalServerError, nil)
		return
	}
	task.Title = input.Title
	task.Done = input.Done
	if err := service.Task.Update(&task); err != nil {
		utils.JSONError(w, "Failed to update task", http.StatusInternalServerError, nil)
		return
	}
	utils.JSONSuccess(w, task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.ParseUint(id, 10, 64)
	// idInt, _ := strconv.Atoi(id)

	// for i, task := range models.Tasks {
	// 	if task.ID == idInt {
	// 		models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
	// 		w.WriteHeader(http.StatusNoContent)
	// 		return
	// 	}
	// }
	// http.Error(w, "Task not found", http.StatusNotFound)
	if err := service.Task.Delete(uint(idInt)); err != nil {
		utils.JSONError(w, "Failed to delete task", http.StatusInternalServerError, nil)
		return
	}
	utils.JSONSuccess(w, nil)
	// var task models.Task
	// if result := database.DB.First(task, id); result.Error != nil {
	// 	http.Error(w, "Tasks not found", http.StatusNotFound)
	// 	return
	// }
	// database.DB.Delete(&task)
	// w.WriteHeader(http.StatusNoContent)
}
