package models

import "time"

type Task struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title" validate:"required,min=2,max=100"`
	Done      bool      `gorm:"default:false" json:"done"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

// var Tasks = []Task{
// 	{ID: 1, Title: "Learn Go", Done: true, CreatedAt: time.Now()},
// 	{ID: 2, Title: "Build REST API", Done: false, CreatedAt: time.Now()},
// 	{ID: 3, Title: "Learn React", Done: false, CreatedAt: time.Now()},
// }

// var NextID = 4
