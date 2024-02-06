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
	DB, err = gorm.Open(sqlite.Open(config.SQLitePath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
