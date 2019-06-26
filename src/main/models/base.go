package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB

// automatically get called by Go
func init() {	
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	conn, e := gorm.Open("sqlite3", os.Getenv("db_path"))
	if e != nil {
		fmt.Print(e)
	}
	db = conn
	db.Debug().AutoMigrate(&User{})
}

func GetDB() *gorm.DB {
	return db
}