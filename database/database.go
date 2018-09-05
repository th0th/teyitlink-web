package database

import (
	"fmt"
	"os"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "link:root@/teyitlink?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred while initiating database connection: %v\n", err)
		os.Exit(1)
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

func Migrate(db *gorm.DB) {
	log.Print("Migrating the database...")
	db.AutoMigrate(&Archive{})
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
