package dto

import "github.com/google/uuid"

type CreateUserDto struct {
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name"`
	SecondName *string `json:"second_name"`
}

type CreateUserResponseDto struct {
	UUID       uuid.UUID `json:"uuid"`
	FirstName  *string   `json:"first_name"`
	SecondName *string   `json:"second_name"`
	Email      string    `json:"email"`
}
