package repositories

import "project/domain/models"

type RefreshTokenRepository interface {
	Add(model *models.RefreshToken) error
}
