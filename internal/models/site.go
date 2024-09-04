package models

import (
	"errors"
	"gorm.io/gorm"
)

type Site struct {
	ID               uint    `json:"id" gorm:"primaryKey"`
	CaptchaSiteKey   string  `json:"captcha_site_key" gorm:"varchar(32)"`
	CaptchaSecretKey string  `json:"captcha_secret_key" gorm:"varchar(32)"`
	DefaultTraffic   float64 `json:"default_traffic"`
}

func (s *Site) BeforeCreate(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Site{}).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("site config already exists")
	}
	return
}
