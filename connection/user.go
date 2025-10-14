package connection

import (
	"gorm.io/gorm"
	"project/domain/models"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Add(model *models.User) error {
	return r.db.Create(model).Error
}

func (r *userRepo) GetByUUID(uuid string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Delete(uuid string) error {
	return r.db.Delete(&models.User{}, "uuid = ?", uuid).Error
}

func (r *userRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
