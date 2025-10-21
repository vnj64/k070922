package connection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project/domain"
	"project/domain/repositories"
	"project/infra/config"
)

type connection struct {
	db *gorm.DB

	userRepository repositories.UserRepository
}

func makeConnection(db *gorm.DB) *connection {
	return &connection{
		db:             db,
		userRepository: NewUserRepository(db),
	}
}

func Make(cfg *config.Config) (domain.Connection, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DbUser(),
		cfg.DbPassword(),
		cfg.DbHost(),
		cfg.DbPort(),
		cfg.DbPassword(),
		cfg.DbSslMode(),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		return nil, fmt.Errorf("unable to open database due [%s]", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to get gorm.DB object due [%s]", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping gorm.DB due [%s]", err)
	}

	return makeConnection(db), nil
}

func (c *connection) UserRepository() repositories.UserRepository {
	return c.userRepository
}
