package models

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            string        `json:"id" gorm:"primaryKey"`
	Username      string        `json:"username" gorm:"not null"`
	Email         string        `json:"email" gorm:"unique;not null"`
	CreatedAt     time.Time     `json:"created_at" gorm:"not null"`
	GoogleProfile GoogleProfile `json:"google_profile"`
}

func (q *User) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == "" {
		var uuid, _ = uuid.NewV7()
		q.ID = uuid.String()
	}
	return
}

type GoogleProfile struct {
	gorm.Model
	ID           string    `json:"google_id" gorm:"primaryKey"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	OAuthToken   string    `json:"oauth_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenExpiry  time.Time `json:"token_expiry"`
	Scopes       []string  `json:"scopes" gorm:"type:text[]"`
	UserID       string    `json:"user_id"`
}

func (q *GoogleProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == "" {
		var uuid, _ = uuid.NewV7()
		q.ID = uuid.String()
	}
	return
}
