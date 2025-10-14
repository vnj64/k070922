package models

import (
	"github.com/google/uuid"
	"time"
)

type RefreshToken struct { // refresh_token
	Token     string    `json:"token" gorm:"primaryKey;"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	Revoked   bool      `json:"revoked"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
