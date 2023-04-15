package repository

import (
	"backend/internal/schema"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(email string) (schema.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByEmail(email string) (schema.User, error) {
	var user schema.User

	err := r.db.Session(&gorm.Session{}).Where("email = ?", email).First(&user).Error

	return user, err
}
