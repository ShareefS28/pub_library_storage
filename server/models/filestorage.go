package models

type Filestorage struct {
	BaseModel

	ID int `gorm:"column:id;primaryKey;autoIncrement"`

	BookID int `gorm:"column:mt_book_id;not null;uniqueIndex"`

	FileName string `gorm:"column:file_name;type:nvarchar(255);not null"`
	FilePath string `gorm:"column:file_path;type:nvarchar(255);not null"`
	FileSize int64  `gorm:"column:file_size;type:bigint;not null"`
	MimeType string `gorm:"column:mime_type;type:nvarchar(255);not null"`
}

func (Filestorage) TableName() string {
	return "mt_file_storage"
}
