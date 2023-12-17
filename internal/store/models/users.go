package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id             uuid.UUID `gorm:"type:uuid;primarykey"`
	Email          string    `gorm:"type:unique"`
	Name           string
	PasswordEncode string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New()

	return
}
