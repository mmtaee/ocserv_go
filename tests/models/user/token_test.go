package user

import (
	"github.com/stretchr/testify/assert"
	"log"
	"ocserv/internal/models"
	"ocserv/pkg/testutils"
	"testing"
)

var (
	dbToken = testutils.GetTestDB()
	userID  uint
)

func init() {
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
