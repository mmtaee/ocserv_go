package user

import (
	"github.com/stretchr/testify/assert"
	"log"
	"ocserv/internal/models"
	"ocserv/pkg/testutils"
	"testing"
)

var dbUser = testutils.GetTestDB()

func init() {
	user := &models.User{
		Username: "test-init",
		Password: "test-passwd",
		IsStaff:  true,
	}
	err := dbUser.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {
	user := &models.User{
		Username: "test",
		Password: "test-passwd",
		IsStaff:  false,
	}
	err := dbUser.Create(&user).Error
	assert.Nil(t, err)
	assert.NotEqual(t, user.ID, uint(0))
}

func TestUpdateUser(t *testing.T) {
	var user models.User

	result := dbUser.First(&user)
	assert.NoError(t, result.Error)

	user.Password = "test-passwd-update"
	result = dbUser.Save(&user)
	assert.NoError(t, result.Error)
	assert.NotEqual(t, user.ID, uint(0))
	assert.Equal(t, user.Password, "test-passwd-update")

	oldUsername := user.Username
	username := "test-init-update"
	result = dbUser.Model(&user).Update("Username", username)
	assert.NoError(t, result.Error)
	assert.NotEqual(t, user.Username, username)
	assert.Equal(t, user.Username, oldUsername)
}

func TestDeleteUser(t *testing.T) {
	var user models.User
	result := dbUser.First(&user)
	assert.NoError(t, result.Error)

	result = dbUser.Delete(&user)
	assert.NoError(t, result.Error)
	assert.NotEqual(t, user.ID, uint(0))
}
