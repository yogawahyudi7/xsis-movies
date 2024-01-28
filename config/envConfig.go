package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type ServerConfig struct {
	Name string
	Port string
	Env  string

	Database struct {
		DBName   string
		Host     string
		Port     string
		Username string
		Password string
		Encrypt  string
		TimeZone string
	}
}

func Get() *ServerConfig {

	err := godotenv.Load(".env")

	if err != nil {
		log.Info("Error loading .env file")
	}

	var defaultConfig ServerConfig

	defaultConfig.Name = os.Getenv("APP_NAME")
	defaultConfig.Port = os.Getenv("APP_PORT")
	defaultConfig.Env = os.Getenv("APP_ENV")

	defaultConfig.Database.TimeZone = os.Getenv("DB_TIME_ZONE")
	defaultConfig.Database.DBName = os.Getenv("DB_NAME")
	defaultConfig.Database.Host = os.Getenv("DB_HOST")
	defaultConfig.Database.Port = os.Getenv("DB_PORT")
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")
	defaultConfig.Database.Encrypt = os.Getenv("DB_ENCRYPT")

	return &defaultConfig
}
