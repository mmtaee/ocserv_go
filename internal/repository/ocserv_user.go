package repository

import (
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/internal/repository/ocserv"
	"ocserv/pkg/database"
)

type OcservUserRepository struct {
	db  *gorm.DB
	cmd ocserv.ServiceOcservRepositoryInterface
}

type OcservUserRepositoryInterface interface {
	GetUserByID(int) (*models.OcservUser, error)
	CreateUser(*models.OcservUser) (*models.OcservUser, error)
	UpdateUser(*models.OcservUser) (*models.OcservUser, error)
	DeleteUser(uint, string) error
}

func NewOcservUserRepository() *OcservUserRepository {
	return &OcservUserRepository{
		db:  database.Connection(),
		cmd: ocserv.NewOcservRepository(),
	}
}

func (o *OcservUserRepository) GetUserByID(ocservUserID int) (*models.OcservUser, error) {
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

func (o *OcservUserRepository) CreateUser(ocservUser *models.OcservUser) (*models.OcservUser, error) {
	ch := make(chan struct {
		ocservUser *models.OcservUser
		err        error
	}, 1)

	err := o.cmd.CreateOrUpdateUser(ocservUser.Group, ocservUser.Username, ocservUser.Password)
	if err != nil {
		return nil, err
	}

	go func() {
		err := o.db.Create(ocservUser).Error
		ch <- struct {
			ocservUser *models.OcservUser
			err        error
		}{ocservUser, err}
	}()
	result := <-ch
	return result.ocservUser, result.err
}

func (o *OcservUserRepository) UpdateUser(ocservUser *models.OcservUser) (*models.OcservUser, error) {
	ch := make(chan struct {
		ocservUser *models.OcservUser
		err        error
	}, 1)

	err := o.cmd.CreateOrUpdateUser(ocservUser.Group, ocservUser.Username, ocservUser.Password)
	if err != nil {
		return nil, err
	}

	go func() {
		err := o.db.Updates(&ocservUser).Error
		ch <- struct {
			ocservUser *models.OcservUser
			err        error
		}{ocservUser, err}
	}()
	result := <-ch
	return result.ocservUser, result.err
}

func (o *OcservUserRepository) DeleteUser(ocservUserID uint, ocservUsername string) error {
	ch := make(chan error, 1)
	err := o.cmd.DeleteUser(ocservUsername)
	if err != nil {
		return err
	}
	go func() {
		ch <- o.db.Delete(&models.OcservUser{}, ocservUserID).Error
	}()
	return <-ch
}
