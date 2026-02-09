package register

import (
	"server/database"
	"server/dtos/dtoregister"
	"server/models"
	"server/utils/utilauthen"
	"server/utils/utilresponse"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// @Tags         Authentication
// @Summary      Create Account
// @Description  Create Account
// @Accept       json
// @Produce      json
// @Param		 request body dtoregister.DTORegisterAccountReq true "Register Account"
// @Router       /auth/register [post]
func RegisterAccount(c fiber.Ctx) error {

	// check body of request
	var req dtoregister.DTORegisterAccountReq
	if err := c.Bind().Body(&req); err != nil {
		return utilresponse.Error(
			c,
			"L00",
			fiber.StatusBadRequest,
			"Invalid request body",
		)
	}

	// chck Account
	var exists bool
	tx_acc := database.DB.Raw(`
		SELECT CASE 
			WHEN EXISTS (
				SELECT 1 
				FROM dbo.mt_accounts 
				WHERE username = ?
			) THEN 1 ELSE 0
		END
	`, req.Username).Scan(&exists)

	if tx_acc.Error != nil {
		return utilresponse.Error(
			c,
			"D00",
			fiber.StatusInternalServerError,
			"Database error",
		)
	}

	if exists {
		return utilresponse.Error(
			c,
			"D00",
			fiber.StatusConflict,
			"Already Has This Username",
		)
	}

	// Create Account
	accUUID := uuid.New()
	hashPassword, err := utilauthen.HashValue(req.Password)
	if err != nil {
		return utilresponse.Error(
			c,
			"Z96",
			fiber.StatusInternalServerError,
			"Error hashing password",
		)
	}

	acc := models.Account{
		BaseModel: models.BaseModel{
			CreatedBy: accUUID,
		},
		UUID:         accUUID,
		Username:     req.Username,
		PasswordHash: hashPassword,
	}

	if err := database.DB.Create(&acc).Error; err != nil {
		return utilresponse.Error(
			c,
			"C10",
			fiber.StatusInternalServerError,
			"Error To Create Account",
		)
	}

	res := dtoregister.DTORegisterAccountSuccessRes{}
	res.User.UUID = acc.UUID
	res.User.Username = acc.Username

	return utilresponse.Success(
		c,
		fiber.StatusOK,
		res,
	)
}
