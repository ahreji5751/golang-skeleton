package models

import "time"

type User struct {
	ID                uint64 `gorm:"AUTO_INCREMENT;primary_key"`
	Email             string `gorm:"type:varchar(100);unique_index;not null" json:"email" create:"nonzero,min=3,max=40,regexp=^[a-z0-9._+]+@[a-z0-9.]+[.][a-z]+$" signin:"nonzero"`
	Name              string `gorm:"not null" json:"name" create:"nonzero,min=3,max=50"`
	EncryptedPassword []byte `gorm:"not null"`
	Tokens            []UserToken
	Password          string `gorm:"-" json:"password" create:"nonzero,min=8" signin:"nonzero"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
