package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	BaseModel

	ID   int       `gorm:"column:id;primaryKey;autoIncrement;not null"`
	UUID uuid.UUID `gorm:"column:uuid;type:uniqueidentifier;not null"`

	AccountID int      `gorm:"column:mt_account_id;not null"`
	Account   *Account `gorm:"foreignKey:AccountID;references:ID;constraint:OnUpdate:CASCADE"`

	IPAddress *string `gorm:"column:ip_address;type:varchar(45)"`
	UserAgent *string `gorm:"column:user_agent;type:varchar(512)"`

	ExpiredAt time.Time  `gorm:"column:expired_at;type:datetime2(7);not null"`
	RevokedAt *time.Time `gorm:"column:revoked_at;type:datetime2(7)"`
}

func (Session) TableName() string {
	return "mt_sessions"
}
