package initializers

import (
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/models"
)

// function runs all migrations and ensures that the database schema is the latest
func SyncDatabase() {
	DB.AutoMigrate(&models.User{},&models.ThreeOddsBet{},&models.TwoOddsBet{})
}