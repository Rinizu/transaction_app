package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret    string
	CustomerFile string
	MerchantFile string
	HistoryFile  string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		JWTSecret:    os.Getenv("JWT_SECRET"),
		CustomerFile: os.Getenv("CUSTOMER_FILE"),
		MerchantFile: os.Getenv("MERCHANT_FILE"),
		HistoryFile:  os.Getenv("HISTORY_FILE"),
	}
}
