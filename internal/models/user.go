package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;type:varchar(32)"`
	IsStaff  bool   `json:"is_staff"`
	Tokens   Token  `json:"token" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
}

type Token struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"user_id"`
	Key      string `json:"key" gorm:"varchar(32)"`
	ExpireAt int64  `json:"expire_at" gorm:"DEFAULT:0"`
}

type OcservUser struct {
	ID             uint            `json:"id" gorm:"primaryKey"`
	Group          string          `json:"group" gorm:"varchar(32)"`
	Username       string          `json:"username" gorm:"unique;varchar(32)"`
	Password       string          `json:"password" gorm:"varchar(64)"`
	IsActive       bool            `json:"is_active"`
	CreatedAt      int64           `json:"created_at"`
	UpdatedAt      int64           `json:"updated_at"`
	ExpireAt       int64           `json:"expire_at" gorm:"DEFAULT:0"`
	RX             float64         `json:"rx"`
	TX             float64         `json:"tx"`
	DefaultTraffic float64         `json:"default_traffic"`
	TrafficType    ServiceTypeEnum `json:"traffic_type" gorm:"varchar(16);default:'FREE'"`
}

func (s *OcservUser) BeforeCreate(tx *gorm.DB) (err error) {
	if s.DefaultTraffic == 0 && s.TrafficType != FREE {
		var config *Site
		err = tx.First(&config).Error
		if err != nil {
			s.DefaultTraffic = 0
		}
		s.DefaultTraffic = config.DefaultTraffic
	}
	if s.TrafficType == FREE {
		s.DefaultTraffic = 0
	}
	return
}

func (s *OcservUser) BeforeUpdate(tx *gorm.DB) (err error) {
	if s.TrafficType != FREE && s.TX > s.DefaultTraffic {
		s.IsActive = false
	}
	if s.TrafficType == FREE {
		s.DefaultTraffic = 0
	}
	return
}
