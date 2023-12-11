package initializers

import (
	"os"

	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/common"
	"gorm.io/gorm"
)

// Gloabl variable that connects to the postgresql database
var DB *gorm.DB

// Connect to database depending on the environment i.e. either test or development.
func ConnectToDatabase(){
	// Check if the environment variable for specifying the environment is set.
	var err error
	environment := os.Getenv("ENVIRONMENT")

	if environment == "test" {
		// If the environment is set to "test," use the test database.
		DB, err = common.TestDBInit()
	} else {
		// For any other environment, use the production database.
		DB, err = common.Init()
	}

	if err != nil {
		panic("failed to connect to the database")
	}
}