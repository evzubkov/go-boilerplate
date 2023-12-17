package users

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository - new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}
