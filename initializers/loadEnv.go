package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

var singletonConfig *Config

type Config struct {
	SQLitePath string `mapstructure:"SQLITE_PATH"`
	ServerPort string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	UploadFolder string `mapstructure:"UPLOAD_FOLDER"`
}

func LoadConfig() (config Config, err error) {
	if singletonConfig != nil {
		return *singletonConfig, nil
	}
	err = godotenv.Load("app.env")
	if err != nil {
		return
	}

	config = Config{
		SQLitePath:   os.Getenv("SQLITE_PATH"),
		ServerPort:   os.Getenv("PORT"),
		ClientOrigin: os.Getenv("CLIENT_ORIGIN"),
		UploadFolder: os.Getenv("UPLOAD_FOLDER"),
	}

	singletonConfig = &config

	return config, nil
}
