package initializers

import "go-fiber-postgres/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
