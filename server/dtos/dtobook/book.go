package dtobook

type DTOCreateBookReq struct {
	Name string `form:"name" validate:"required"`
}

type DTOCreateBookSuccessReq struct {
	Name string `json:"name"`
}
