package models

import (
	"time"

	"github.com/google/uuid"
)

// Customer struct to describe User object.
type Customer struct {
	ID          uuid.UUID `gorm:"primaryKey;not null;unique" db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time `gorm:"not null" db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" db:"updated_at" json:"updated_at"`
	FirstName   string    `gorm:"not null" db:"firstName" json:"firstName" validate:"required,lte=25"`
	LastName    string    `gorm:"not null" db:"lastName" json:"lastName" validate:"required,email,lte=255"`
	DOB         time.Time `gorm:"not null" db:"dob" json:"dob" validate:"required,lte=255"`
	Address     string    `gorm:"not null" db:"address" json:"address" validate:"required,len=1"`
	City        string    `gorm:"not null" db:"city" json:"city" validate:"required,lte=6"`
	PostalCode  int       `gorm:"notn ull" db:"postalCode" json:"postalCode" validate:"required,lte=9"`
	NumberPhone int       `gorm:"not null" db:"number_phone" json:"number_phone" validate:"required,lte=15"`
}
