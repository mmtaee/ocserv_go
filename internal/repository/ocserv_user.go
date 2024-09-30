package repository

import (
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/database"
)

type OcservRepository struct {
	db *gorm.DB
}

type OcservRepositoryInterface interface {
	GetUserByID(int) (*models.OcservUser, error)
	CreateUser(*models.OcservUser) (*models.OcservUser, error)
	UpdateUser(*models.OcservUser) (*models.OcservUser, error)
	DeleteUser(uint) error
}

func NewOcservRepository() *OcservRepository {
	return &OcservRepository{
		db: database.Connection(),
	}
}

func (o *OcservRepository) GetUserByID(ocservUserID int) (*models.OcservUser, error) {
	ch := make(chan struct {
		ocservUser *models.OcservUser
		err        error
	}, 1)

	go func() {
		var ocservUser *models.OcservUser
		err := o.db.Where("id = ?", ocservUserID).First(&ocservUser).Error
		ch <- struct {
			ocservUser *models.OcservUser
			err        error
		}{ocservUser, err}
	}()
	result := <-ch
	return result.ocservUser, result.err
}

func (o *OcservRepository) CreateUser(ocservUser *models.OcservUser) (*models.OcservUser, error) {
	ch := make(chan struct {
		ocservUser *models.OcservUser
		err        error
	}, 1)

	go func() {
		err := o.db.Create(ocservUser).Error
		ch <- struct {
			ocservUser *models.OcservUser
			err        error
		}{ocservUser, err}
	}()

	//	TODO: call ocserv service
	result := <-ch
	return result.ocservUser, result.err
}

func (o *OcservRepository) UpdateUser(ocservUser *models.OcservUser) (*models.OcservUser, error) {
	ch := make(chan struct {
		ocservUser *models.OcservUser
		err        error
	}, 1)

	go func() {
		err := o.db.Updates(&ocservUser).Error
		ch <- struct {
			ocservUser *models.OcservUser
			err        error
		}{ocservUser, err}
	}()

	//	TODO: call ocserv service
	result := <-ch
	return result.ocservUser, result.err
}

func (o *OcservRepository) DeleteUser(ocservUserID uint) error {
	ch := make(chan error, 1)

	go func() {
		ch <- o.db.Delete(&models.OcservUser{}, ocservUserID).Error
	}()
	return <-ch
}
