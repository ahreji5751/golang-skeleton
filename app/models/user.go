package models

import "time" // if you need/want

type User struct { // example user fields
	ID                int64  `gorm:"AUTO_INCREMENT;primary_key"`
	Email             string `gorm:"type:varchar(100);unique_index" json:"email" validate:"min=3,max=40,regexp=^[a-z0-9._+]+@[a-z0-9.]+[.][a-z]+$"`
	EncryptedPassword []byte `gorm:"index"`
	Token             string `gorm:"type:varchar(80);index"`
	Password          string `gorm:"-" json:"password" validate:"min=8"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
