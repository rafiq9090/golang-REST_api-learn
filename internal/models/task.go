package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

var Tasks = []Task{
	{ID: 1, Title: "Learn Go", Done: true, CreatedAt: time.Now()},
	{ID: 2, Title: "Build REST API", Done: false, CreatedAt: time.Now()},
	{ID: 3, Title: "Learn React", Done: false, CreatedAt: time.Now()},
}

var NextID = 4
