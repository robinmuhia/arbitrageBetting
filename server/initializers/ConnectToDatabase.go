package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(){
	var err error
	dsn := os.Getenv("DB_CONNECTION")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic("failed to connect to postgre database")
	}
}