package migrations

import (
	"fiber/config"
	"fiber_project/models"
	"log"
)

func UserSeeder() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
     
	// If users already exist, skip seeding
	if count > 0 {
		log.Println("✅ Users already exist. Skipping seeding.")
		return
	}
	// Step 2: Define users to seed
	// users := []models.User{
	// 	{Name: "Anand", Email: "anand@example.com", Password: "password123"},
	// 	{Name: "Ravi", Email: "ravi@example.com", Password: "password456"},
	// 	{Name: "Sita", Email: "sita@example.com", Password: "password789"},
	// }

	users := []models.User{
		{name : "Anand"},
		{name : "Shiva"},
		{name : "Kambo"},
	}

	// Step 3: Insert users into the database
	// for _, user := range users {
	// 	if err := config.DB.Create(&user).Error; err != nil {
	// 		log.Fatalf("❌ Failed to seed user: %v", err)
	// 	}
	// }

	for i := 0; i < len(users); i++ {
		user := users[i]
		if err := config.DB.Create(&user).Error; err != nil {
			log.Fatalf("❌ Failed to seed user: %v", err)
		}
	}


	log.Println("✅ Users seeded successfully!")
}
