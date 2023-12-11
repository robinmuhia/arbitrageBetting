package common

import (
	"fmt"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database is a struct that embeds a *gorm.DB instance, representing a connection to a database.
type Database struct {
	*gorm.DB
}

// DB is a global variable that holds the main *gorm.DB instance for the application.
var DB *gorm.DB

// Opens a database and save the reference to `Database` struct.
func Init() (*gorm.DB, error) {
    dsn := os.Getenv("DB_CONNECTION")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("db err: (Init)", err)
        return nil, err
    }
	sqlDB, _ := db.DB()  // Ignore the error
    sqlDB.SetMaxIdleConns(10)
    // db.LogMode(true)
    DB = db
    return DB, nil
}



// This function will create a temporarily database for running testing cases
func TestDBInit() (*gorm.DB, error) {
	dsn := "host=localhost user=tests password=tests123 dbname=testarbitragedb port=5432 sslmode=disable"
	test_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("db err: (Init)", err)
        return nil, err
    }
	testSqlDb,_ := test_db.DB()
	testSqlDb.SetMaxIdleConns(3)
	DB = test_db
	return DB,nil
}


// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}