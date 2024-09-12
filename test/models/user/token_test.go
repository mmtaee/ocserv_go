package user

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"ocserv/internal/models"
	"ocserv/pkg/testutils"
	"testing"
	"time"
)

var (
	dbToken  *gorm.DB
	userID   uint
	expireAt = time.Now().Add(time.Hour * 2).Unix()
	tokenID  uint
)

func init() {
	dbToken = testutils.GetTestDB()
	user := &models.User{
		Username: "test-init-token",
		Password: "test-passwd",
		IsStaff:  true,
	}
	err := dbToken.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	userID = user.ID

	token := &models.Token{
		UserID:   userID,
		Key:      "test-key-init",
		ExpireAt: expireAt,
	}
	err = dbToken.Create(token).Error
	if err != nil {
		log.Fatal(err)
	}
	tokenID = token.ID
}

func TestCreateToken(t *testing.T) {
	expireAtCreate := time.Now().Add(time.Hour).Unix()
	token := &models.Token{
		UserID:   userID,
		Key:      "test-key",
		ExpireAt: expireAtCreate,
	}
	err := dbToken.Create(token).Error
	assert.Nil(t, err)
	assert.NotEqual(t, token.ID, uint(0))
	assert.Equal(t, token.ExpireAt, expireAtCreate)
	assert.Equal(t, token.Key, "test-key")
}

func TestUpdateToken(t *testing.T) {
	var token models.Token
	err := dbToken.First(&token, userID).Error
	assert.Nil(t, err)
	assert.NotEqual(t, token.ID, uint(0))
	assert.Equal(t, token.ExpireAt, expireAt)

	expireAtUpdate := time.Now().Add(time.Hour * 5).Unix()

	token.ExpireAt = expireAtUpdate
	token.Key = "test-key-update"
	err = dbToken.Save(&token).Error
	assert.Nil(t, err)

	err = dbToken.First(&token, userID).Error
	assert.NotEqual(t, token.ID, uint(0))
	assert.Equal(t, token.ExpireAt, expireAtUpdate)
	assert.NotEqual(t, token.Key, "test-key-update")
}

func TestDeleteToken(t *testing.T) {
	var token models.Token

	err := dbToken.Delete(&models.Token{}, tokenID).Error
	assert.Nil(t, err)

	err = dbToken.First(&token, tokenID).Error
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
