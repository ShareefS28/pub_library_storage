package models

type Book struct {
	BaseModel

	ID int `gorm:"column:id;primaryKey;autoIncrement"`

	AccountID int      `gorm:"column:mt_account_id;not null"`
	Account   *Account `gorm:"constraint:foreignKey:AccountID;references:ID;OnDelete:NO ACTION;"`

	Name string `gorm:"column:name;type:nvarchar(255);not null"`

	Filestorage Filestorage `gorm:"foreignKey:BookID;references:ID"`
}

func (Book) TableName() string {
	return "mt_books"
}
