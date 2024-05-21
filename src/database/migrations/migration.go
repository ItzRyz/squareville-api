package migrations

import (
	"fmt"
	"log"

	"github.com/ItzRyz/squareville-api/src/database"
	"github.com/ItzRyz/squareville-api/src/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}
