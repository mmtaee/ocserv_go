package repository

import (
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/database"
)

type TokenRepository struct {
	db *gorm.DB
}
type TokenRepositoryInterface interface {
	GetTokenByKey(string) (*models.User, *models.Token, error)
	Create(*models.Token) (*models.Token, error)
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		db: database.Connection(),
	}
}

func (t *TokenRepository) GetTokenByKey(key string) (*models.User, *models.Token, error) {
	ch := make(chan struct {
		user  *models.User
		token *models.Token
		err   error
	}, 1)

	go func() {
		var token *models.Token
		err := t.db.Where("key = ?", key).First(&token).Error

		if err != nil {
			ch <- struct {
				user  *models.User
				token *models.Token
				err   error
			}{nil, nil, err}
			return
		}

		var user *models.User
		err = t.db.Where("id = ? ", token.UserID).First(&user).Error
		if err != nil {
			ch <- struct {
				user  *models.User
				token *models.Token
				err   error
			}{nil, nil, err}
			return
		}

		ch <- struct {
			user  *models.User
			token *models.Token
			err   error
		}{user, token, err}
	}()
	result := <-ch
	return result.user, result.token, result.err
}

func (t *TokenRepository) Create(token *models.Token) (*models.Token, error) {
	ch := make(chan struct {
		token *models.Token
		err   error
	})

	go func() {
		err := t.db.Create(&token).Error
		if err != nil {
			ch <- struct {
				token *models.Token
				err   error
			}{nil, err}
			return
		}
		ch <- struct {
			token *models.Token
			err   error
		}{token, err}
	}()
	result := <-ch
	return result.token, result.err
}
