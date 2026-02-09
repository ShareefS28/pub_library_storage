package dtoprofile

import (
	"time"

	"github.com/google/uuid"
)

type DTOProfileSuccessRes struct {
	User struct {
		UUID      uuid.UUID `json:"uuid"`
		ExpiredAt time.Time `json:"expired_at"`
	} `json:"user"`
}
