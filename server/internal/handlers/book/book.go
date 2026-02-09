package book

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"server/database"
	"server/dtos/dtobook"
	"server/models"
	"server/utils/utilresponse"
	"server/utils/utilsession"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// Max File Size
const maxFileSize = 5 * 1024 * 1024 // 5MB

// Allowed MIME types
var allowedMimeTypes = map[string]bool{
	"application/pdf": true,
	"image/jpeg":      true,
	"image/png":       true,
}

// Allowed extensions
var allowedExtensions = map[string]bool{
	".pdf":  true,
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

// @Tags         Book
// @Summary      Create Book
// @Description  Example:
// @Description  names=Book A
// @Description  names=Book B
// @Description  files=@file1.pdf
// @Description  files=@file2.pdf
// @Accept       multipart/form-data
// @Produce      multipart/form-data
// @Param		 names formData []string true "Book names" collectionFormat(multi)
// @Param		 files formData []file   true "Book files" collectionFormat(multi)
// @Router       /secure/book/create [post]
func CreateBook(c fiber.Ctx) error {

	session, err := utilsession.GetSessionInfo(c)
	if err != nil {
		return utilresponse.Error(c, "L99", fiber.StatusInternalServerError, err.Error())
	}

	// check Content-Type
	form, err := c.MultipartForm()
	if err != nil {
		return utilresponse.Error(
			c,
			"N99",
			fiber.StatusBadRequest,
			"Invalid Content-Type",
		)
	}

	// Get File
	names := form.Value["names"]
	files := form.File["files"]

	if len(names) == 0 {
		return utilresponse.Error(
			c,
			"B01",
			fiber.StatusBadRequest,
			"Names is required",
		)
	}

	if len(files) == 0 {
		return utilresponse.Error(
			c,
			"F99",
			fiber.StatusBadRequest,
			"File is required",
		)
	}

	// limit: max number of files
	if len(names) != len(files) {
		return utilresponse.Error(
			c,
			"F98",
			fiber.StatusBadRequest,
			"Names and Files are not matching",
		)
	}

	if len(files) > 5 {
		return utilresponse.Error(
			c,
			"F98",
			fiber.StatusBadRequest,
			"Max Files Is 5 file",
		)
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return utilresponse.Error(
			c,
			"F10",
			fiber.StatusInternalServerError,
			"Cannot create upload directory",
		)
	}

	// start connection to DB
	conn := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			conn.Rollback()
		}
	}()

	res := make([]dtobook.DTOCreateBookSuccessReq, 0, len(files))

	// Loops files matching
	for i := 0; i < len(files); i++ {

		bookReq := strings.TrimSpace(names[i])
		fileHeader := files[i]

		// book validation
		if strings.TrimSpace(bookReq) == "" {
			conn.Rollback()
			return utilresponse.Error(
				c,
				"B03",
				fiber.StatusBadRequest,
				fmt.Sprintf("Book name is required at index %d", i),
			)
		}

		// size chunk
		if fileHeader.Size > maxFileSize {
			conn.Rollback()
			return utilresponse.Error(
				c,
				"F97",
				fiber.StatusBadRequest,
				fmt.Sprintf("File size exceed Limit %d > %d", fileHeader.Size, maxFileSize),
			)
		}

		// extension file name check
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if !allowedExtensions[ext] {
			conn.Rollback()
			return utilresponse.Error(
				c,
				"F04",
				fiber.StatusBadRequest,
				"File extension not allowed",
			)
		}

		// open file
		file, err := fileHeader.Open()
		if err != nil {
			return utilresponse.Error(
				c,
				"F05",
				fiber.StatusInternalServerError,
				"Cannot open file",
			)
		}

		// Read first 512 bytes for MIME
		buffer := make([]byte, 512)
		n, err := file.Read(buffer)
		if err != nil {
			file.Close()
			return utilresponse.Error(
				c,
				"F06",
				fiber.StatusInternalServerError,
				"Cannot read file",
			)
		}

		mimeType := http.DetectContentType(buffer[:n])
		if !allowedMimeTypes[mimeType] {
			file.Close()
			return utilresponse.Error(
				c,
				"F07",
				fiber.StatusBadRequest,
				"Invalid file content",
			)
		}

		// Reset file pointer
		if _, err := file.Seek(0, 0); err != nil {
			file.Close()
			return utilresponse.Error(
				c,
				"F08",
				fiber.StatusInternalServerError,
				"Cannot process file",
			)
		}

		file.Close()

		// storage file
		storedFileName := uuid.New().String() + ext
		savePath := filepath.Join(uploadDir, storedFileName)

		// save file
		if err := c.SaveFile(fileHeader, savePath); err != nil {
			conn.Rollback()
			return utilresponse.Error(
				c,
				"F09",
				fiber.StatusInternalServerError,
				"Failed to save file",
			)
		}

		// save to DB
		book := models.Book{
			Name:      bookReq,
			AccountID: session.AccountID,
			BaseModel: models.BaseModel{
				CreatedBy: session.AccountUUID,
			},
			Filestorage: models.Filestorage{
				BaseModel: models.BaseModel{
					CreatedBy: session.AccountUUID,
				},
				FileName: fileHeader.Filename,
				FilePath: savePath,
				FileSize: fileHeader.Size,
				MimeType: mimeType,
			},
		}

		if err := conn.Create(&book).Error; err != nil {
			os.Remove(savePath)
			conn.Rollback()
			return utilresponse.Error(
				c,
				"D01",
				fiber.StatusInternalServerError,
				"Database Error",
			)
		}

		res = append(res, dtobook.DTOCreateBookSuccessReq{
			Name: fileHeader.Filename,
		})
	}

	if err := conn.Commit().Error; err != nil {
		conn.Rollback()
		return utilresponse.Error(
			c,
			"T99",
			fiber.StatusInternalServerError,
			"Transaction commit failed",
		)
	}

	return utilresponse.Success(
		c,
		fiber.StatusOK,
		res,
	)
}
