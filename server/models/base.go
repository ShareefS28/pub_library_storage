package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime2(7);not null;default:SYSUTCDATETIME()"`
	CreatedBy uuid.UUID  `gorm:"column:created_by;type:uniqueidentifier;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime2(7);autoUpdateTime:false"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by;type:uniqueidentifier"`
	IsUpdated bool       `gorm:"column:is_updated;type:bit;not null;default:false"`

	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime2(7)"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by;type:uniqueidentifier"`
	IsDeleted bool       `gorm:"column:is_deleted;type:bit;not null;default:false"`
}
