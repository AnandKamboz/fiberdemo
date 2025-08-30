package migrations

import (
	"log"
	"fiber/config"
)

func RemoveCityColumn() {
	err := config.DB.Exec("ALTER TABLE users DROP COLUMN city").Error
	if err != nil {
		log.Fatal("❌ Failed to remove column 'city':", err)
	}
	log.Println("✅ Column 'city' removed successfully!")
}
