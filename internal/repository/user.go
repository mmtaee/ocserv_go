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
	GetUserById(int) (*models.User, error)
	UpdatePassword(uint, string) error
	Exists() (bool, error)
	DeleteStaffUserByID(int) error
	GetUserByUsername(string) (*models.User, error)
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

func (u *UserRepository) GetUserById(id int) (*models.User, error) {
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

func (u *UserRepository) Exists() (bool, error) {
	ch := make(chan struct {
		ok  bool
		err error
	})

	go func() {
		var count int64
		err := u.db.Model(&models.User{}).Count(&count).Error
		if err != nil {
			ch <- struct {
				ok  bool
				err error
			}{false, err}
			return
		}
		ch <- struct {
			ok  bool
			err error
		}{count > 0, err}
	}()
	result := <-ch
	return result.ok, result.err
}

func (u *UserRepository) DeleteStaffUserByID(id int) error {
	ch := make(chan error)
	go func() {
		err := u.db.Where("id = ? AND is_admin = ?", id, false).Delete(&models.User{}).Error
		ch <- err
	}()
	return <-ch
}

func (u *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
