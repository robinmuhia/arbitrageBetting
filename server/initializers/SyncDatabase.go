package initializers

import (
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{},&models.ThreeOddsBet{},&models.TwoOddsBet{})
}