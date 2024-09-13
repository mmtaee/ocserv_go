package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"type:varchar(32);unique;not null;<-:create"`
	Password string `json:"password" gorm:"varchar(64);not null"`
	IsStaff  bool   `json:"is_staff"`
	Tokens   Token  `json:"token" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
}

type Token struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"user_id"`
	Key      string `json:"key" gorm:"varchar(32);not null;unique;<-:create"`
	ExpireAt int64  `json:"expire_at" gorm:"DEFAULT:0"`
}
