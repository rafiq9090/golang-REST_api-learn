package repository

import (
	"task-api/internal/database"
	"task-api/internal/models"
)

type TaskRepo struct{}

var Task = TaskRepo{}

func (TaskRepo) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	result := database.DB.Find(&tasks)
	return tasks, result.Error
}

func (TaskRepo) GetByID(id uint) (models.Task, error) {
	var task models.Task
	err := database.DB.First(&task, id).Error
	return task, err
}

func (TaskRepo) Create(task *models.Task) error {
	return database.DB.Create(task).Error
}

func (TaskRepo) Update(task *models.Task) error {
	return database.DB.Save(task).Error
}

// func (TaskRepo) Delete(task *models.Task) error {
// 	return database.DB.Delete(&models.Task{}, task.ID).Error
// }

func (TaskRepo) Delete(id uint) error {
	return database.DB.Delete(&models.Task{}, id).Error
}
