package initializers

import (
	"fmt"
	"log"

	// "gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	log.Printf("Connecting to the Database: %s", config.SQLitePath)
	DB, err = gorm.Open(sqlite.Open(config.SQLitePath+"?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	if res := DB.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		log.Fatal("Failed to enable foreign keys")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
