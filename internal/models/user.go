// User represents the user model.
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;not null" binding:"required"`
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Age       int
}
