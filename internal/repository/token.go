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
		err := t.db.Where("key = ?", key).First(&token).Joins("User").Error

		if err != nil {
			ch <- struct {
				user  *models.User
				token *models.Token
				err   error
			}{nil, nil, err}
			return
		}

		var user *models.User
		err = t.db.First(&user, token.UserID).Error
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
