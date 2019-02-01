package models

import "time"

type UserToken struct {
	ID        uint64 `gorm:"AUTO_INCREMENT;primary_key"`
	Token     string `gorm:"type:varchar(64);index;not null"`
	Active    uint   `gorm:"type:tinyint(1);default: 1" json:"active"`
	UserID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
