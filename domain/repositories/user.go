package repositories

import "project/domain/models"

type UserRepository interface {
	Add(model *models.User) error
	GetByUUID(uuid string) (*models.User, error)
	Delete(uuid string) error
	GetByEmail(email string) (*models.User, error)
}
