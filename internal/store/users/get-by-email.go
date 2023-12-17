package users

import "boilerplate/internal/store/models"

// GetByEmail - get user by email
func (o *UserRepository) GetByEmail(email string) (user models.User, err error) {

	user = models.User{Email: email}
	result := o.db.Find(&user)

	return user, result.Error
}
