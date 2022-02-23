package config

import (
	"github.com/ToxicityMax/smol/logger"
	"github.com/joho/godotenv"
	"os"
)

type config struct {
	SECRET string `default:"secret-key"`
	DEBUG  string `default:"release"`
}

var C config

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		logger.Error("Error loading .env file")
	} else {
		C = config{
			SECRET: os.Getenv("SECRET"),
			DEBUG:  os.Getenv("DEBUG"),
		}
	}
}
