package dtoregister

import "github.com/google/uuid"

type DTORegisterAccountReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type DTORegisterAccountSuccessRes struct {
	User struct {
		UUID     uuid.UUID `json:"uuid"`
		Username string    `json:"username"`
	} `json:"user"`
}
