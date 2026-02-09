package models

import (
	"github.com/google/uuid"
)

type Account struct {
	BaseModel

	ID   int       `gorm:"column:id;primaryKey;autoIncrement"`
	UUID uuid.UUID `gorm:"column:uuid;type:uniqueidentifier;not null"`

	Username     string `gorm:"column:username;type:varchar(255);unique;not null"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(512);not null"`

	Sessions []Session `gorm:"foreignKey:AccountID;references:ID;constraint:OnUpdate:CASCADE;"`
	Books    []Book    `gorm:"foreignKey:AccountID;references:ID;"`
}

func (Account) TableName() string {
	return "mt_accounts"
}
