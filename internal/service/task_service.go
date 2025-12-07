package service

import (
	"task-api/internal/models"
	"task-api/internal/repository"
)

type TaskService struct{}

var Task = TaskService{}

func (TaskService) GetAll() ([]models.Task, error) {
	return repository.Task.GetAll()
}

func (TaskService) GetByID(id uint) (models.Task, error) {
	return repository.Task.GetByID(id)
}

func (TaskService) Create(title string, done bool) (models.Task, error) {
	task := models.Task{
		Title: title,
		Done:  done,
	}
	err := repository.Task.Create(&task)
	return task, err
}

func (TaskService) Update(task *models.Task) error {
	return repository.Task.Update(task)
}

func (TaskService) Delete(id uint) error {
	return repository.Task.Delete(id)
}
