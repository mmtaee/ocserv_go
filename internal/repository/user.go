package repository

import (
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/database"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	Create(*models.User) (*models.User, error)
	GetUserById(int64) (*models.User, error)
	UpdatePassword(uint, string) error
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.Connection(),
	}
}

func (u *UserRepository) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(user).Error
	return user, err
}

func (u *UserRepository) GetUserById(id int64) (*models.User, error) {
	var user models.User
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) UpdatePassword(userID uint, password string) error {
	ch := make(chan error)
	go func() {
		err := u.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
		ch <- err
	}()
	return <-ch
}
