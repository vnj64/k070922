package domain

import "project/domain/repositories"

type Connection interface {
	UserRepository() repositories.UserRepository
}
