package profile

import (
	"server/dtos/dtoprofile"
	"server/utils/utilresponse"
	"server/utils/utilsession"

	"github.com/gofiber/fiber/v3"
)

// @Tags         Profile
// @Summary      Get Profile
// @Description  Get Profile or Check Sessions
// @Accept       json
// @Produce      json
// @Router       /secure/me [get]
func Me(c fiber.Ctx) error {

	session, err := utilsession.GetSessionInfo(c)
	if err != nil {
		return utilresponse.Error(c, "L99", fiber.StatusInternalServerError, err.Error())
	}

	res := dtoprofile.DTOProfileSuccessRes{}
	res.User.UUID = session.AccountUUID
	res.User.ExpiredAt = session.ExpiredAt

	return utilresponse.Success(
		c,
		fiber.StatusOK,
		res,
	)
}
