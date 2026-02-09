package dtologin

import (
	"time"

	"github.com/google/uuid"
)

type DTOLoginAccountReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type DTOLoginSuccessRes struct {
	User struct {
		UUID      uuid.UUID `json:"uuid"`
		Username  string    `json:"username"`
		ExpiredAt time.Time `json:"expired_at"`
	} `json:"user"`
}
