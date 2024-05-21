package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" form:"username" validate:"required" gorm:"not null"`
	Email     string         `json:"email" form:"email" validate:"required,email" gorm:"not null"`
	Password  string         `json:"password" form:"password" validate:"required" gorm:"not null,colum:password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
