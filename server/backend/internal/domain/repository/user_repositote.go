package repository

import (
	"server/internal/domain/models"

	"gorm.io/gorm"
)

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *models.User) error{
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByLogin(login string) (*models.User, error){
	var user models.User
	if err := r.db.Where("login=?", login).First(&user).Error; err != nil{
		return nil, err
	}
	return &user, nil
}