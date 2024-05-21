package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	User      User           `gorm:"foreignKey:ID"`
	Token     string         `json:"token" form:"token" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	ExpiredAt time.Time      `json:"expired_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Login struct {
	Username string `json:"username" form:"username" validate:"required" gorm:"not null"`
	Password string `json:"password" form:"password" validate:"required"`
}
