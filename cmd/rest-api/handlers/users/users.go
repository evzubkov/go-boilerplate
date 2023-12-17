package users

import (
	"boilerplate/internal/config"
	"gorm.io/gorm"
)

type Handler struct {
	db     *gorm.DB
	config config.TomlConfig
}

func NewUsersHandler(db *gorm.DB, config config.TomlConfig) Handler {
	return Handler{
		db:     db,
		config: config,
	}
}
