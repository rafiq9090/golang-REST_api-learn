package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name" validate:"required,min=2,max=100"`
	Description string    `gorm:"type:text" json:"description" validate:"required"`
	Price       uint      `gorm:"not null" json:"price" validate:"required"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
