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
	dbOcservUser *gorm.DB
	ocservUserID uint
)

func init() {
	dbOcservUser = testutils.GetTestDB()
	ocservUser := models.OcservUser{
		Username:       "username",
		Password:       "password",
		ExpireAt:       time.Now().Add(24 * time.Hour).Unix(),
		Group:          "defaults",
		DefaultTraffic: 10,
		IsActive:       true,
		TrafficType:    models.MONTHLY,
	}
	err := dbOcservUser.Create(&ocservUser).Error
	if err != nil {
		log.Fatal(err)
	}
	ocservUserID = ocservUser.ID
}

func TestCreateOcservUser(t *testing.T) {
	ocservUser := models.OcservUser{
		Username:       "username-create",
		Password:       "password-create",
		ExpireAt:       time.Now().Add(24 * time.Hour).Unix(),
		Group:          "defaults",
		DefaultTraffic: 20,
		IsActive:       true,
		TrafficType:    models.MONTHLY,
	}
	err := dbOcservUser.Create(&ocservUser).Error
	assert.Nil(t, err)
	assert.Equal(t, ocservUser.DefaultTraffic, float64(20))
	assert.Equal(t, ocservUser.TrafficType, models.MONTHLY)
	assert.Equal(t, ocservUser.IsActive, true)
}

func TestUpdateOcservUser(t *testing.T) {
	var ocservUser models.OcservUser
	err := dbOcservUser.First(&ocservUser, ocservUserID).Error
	assert.Nil(t, err)

	ocservUser.Username = "username-update"
	ocservUser.Password = "password-update"
	ocservUser.ExpireAt = time.Now().Add(4 * 24 * time.Hour).Unix()
	ocservUser.TrafficType = models.FREE
	err = dbOcservUser.Save(&ocservUser).Error
	assert.Nil(t, err)

	err = dbOcservUser.First(&ocservUser, ocservUserID).Error
	assert.Nil(t, err)
	assert.Equal(t, ocservUser.TrafficType, models.FREE)
	assert.Equal(t, ocservUser.Username, "username")

}

func TestDeleteOcservUser(t *testing.T) {
	var ocservUser models.OcservUser

	err := dbOcservUser.Delete(&models.OcservUser{}, ocservUserID).Error
	assert.Nil(t, err)

	err = dbOcservUser.First(&ocservUser, ocservUserID).Error
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
