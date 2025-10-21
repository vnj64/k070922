package models

import (
	"fmt"
	"github.com/google/uuid"
	"project/pkg/permissions"
	"project/pkg/timestamps"
)

type User struct {
	Timestamps timestamps.Timestamps `json:"timestamps" gorm:"-"`
	UUID       uuid.UUID             `json:"uuid"`
	FirstName  *string               `json:"first_name"`
	SecondName *string               `json:"second_name"`
	Email      string                `json:"email"`
	Role       permissions.Role      `json:"role"`
}

func (u User) Validation() error {
	if u.FirstName != nil && len(*u.FirstName) > 0 {
		return nil
	}
	return fmt.Errorf("first name is incorrect")
}
