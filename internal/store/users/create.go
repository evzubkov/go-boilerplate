package users

import (
	"boilerplate/internal/store/models"
)

// Create - create new user
func (o *UserRepository) Create(user *models.User) (err error) {
	result := o.db.Create(&user)

	return result.Error
}
