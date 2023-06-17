package models

import (
	"time"

	"github.com/google/uuid"
)

// Customer struct to describe User object.
type Account struct {
	ID            uuid.UUID `gorm:"primaryKey;not null;unique" db:"id" json:"id" validate:"required,uuid"`
	CreatedAt     time.Time `gorm:"not null" db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"not null" db:"updated_at" json:"updated_at"`
	Disabled      bool      `gorm:"notn ull" db:"postalCode" json:"postalCode" validate:"required,lte=9"` // checked before all action on the account if true can do nothing
	Type          int       `gorm:"not null" db:"type" json:"type" validate:"required,lte=25"`
	BankAccountID string    `gorm:"not null" db:"bank_account_id" json:"bank_account_id" validate:"required,email,lte=25"`
	UserId        string    `gorm:"not null" db:"user_id" json:"user_id" validate:"required,lte=50"`
	Country       string    `gorm:"not null" db:"country" json:"country" validate:"required,lte=80"`
	// paramater credit, max spending ...
	Balance float32 `gorm:"not null" db:"city" json:"city" validate:"required,lte=35"`
}
