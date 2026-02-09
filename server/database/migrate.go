package database

import (
	"log"
	"server/models"
)

func Migrate() {
	mg := DB.AutoMigrate(
		&models.Account{},
		&models.Session{},
		&models.Book{},
		&models.Filestorage{},
	)

	// migrations.CreateSessionCleanUpJob(DB)
	// migrations.DropSessionCleanupJob(DB)

	if mg != nil {
		log.Println("Migrate Failed!")
	} else {
		log.Println("Migrate Success!")
	}

}
