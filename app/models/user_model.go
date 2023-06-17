package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type User struct {
	ID           uuid.UUID `gorm:"primaryKey;not null;unique" db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `gorm:"not null" db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null" db:"updated_at" json:"updated_at"`
	Username     string    `gorm:"not null" db:"username" json:"username" validate:"required,lte=25"`
	Email        string    `gorm:"not null" db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `gorm:"not null" db:"password_hash" json:"password_hash" validate:"required,lte=255"`
	UserRole     string    `gorm:"not null" db:"user_role" json:"user_role" validate:"required,lte=6"`
	Disabled     bool      `gorm:"notn ull,default:false"`
}
