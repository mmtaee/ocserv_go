package site

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/testutils"
	"testing"
)

var dbSite *gorm.DB

func init() {
	dbSite = testutils.GetTestDB()
}

func TestCreateSite(t *testing.T) {
	site := &models.Site{
		CaptchaSiteKey:   "",
		CaptchaSecretKey: "",
		DefaultTraffic:   20,
	}
	err := dbSite.Create(site).Error
	assert.Nil(t, err)
	assert.NotEqual(t, site.ID, uint(0))

	site2 := &models.Site{
		CaptchaSiteKey:   "",
		CaptchaSecretKey: "",
		DefaultTraffic:   20,
	}
	err = dbSite.Create(site2).Error
	assert.EqualError(t, err, "site config already exists")
	assert.Equal(t, site2.ID, uint(0))
}

func TestGetSite(t *testing.T) {
	var site models.Site

	result := dbSite.First(&site)
	assert.NoError(t, result.Error)
}

func TestUpdateSite(t *testing.T) {
	var site models.Site

	err := dbSite.First(&site).Error
	assert.NoError(t, err)
	site.CaptchaSiteKey = "site_ket_test"
	err = dbSite.Save(&site).Error
	assert.NoError(t, err)
	assert.NotEmpty(t, site.CaptchaSiteKey)
	assert.Equal(t, site.CaptchaSiteKey, "site_ket_test")
}

func TestDeleteSite(t *testing.T) {
	var (
		site  models.Site
		count int64
	)

	err := dbSite.First(&site).Error
	assert.NoError(t, err)
	err = dbSite.Delete(&site).Error
	assert.NoError(t, err)
	err = dbSite.Model(&site).Count(&count).Error
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}
