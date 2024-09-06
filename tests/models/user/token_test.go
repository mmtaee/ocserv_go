package user

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"ocserv/internal/models"
	"ocserv/pkg/testutils"
	"testing"
)

var (
	dbToken *gorm.DB
	userID  uint
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
}

func TestCreateToken(t *testing.T) {
	token := &models.Token{
		UserID: userID,
	}
	err := dbToken.Create(token).Error
	assert.Nil(t, err)
	assert.NotEqual(t, token.ID, uint(0))
}

func TestUpdateToken(t *testing.T) {}

func TestDeleteToken(t *testing.T) {}
