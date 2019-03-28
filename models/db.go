package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPassword := os.Getenv("db_password")

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, username, dbPassword, dbName) //Build connection string
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}, &Task{}, Project{}) //Database migration
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
