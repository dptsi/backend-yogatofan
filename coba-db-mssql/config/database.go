package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Dbname   string
}

func GetDBConfig() *DbConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	return &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Dbname:   os.Getenv("DB_NAME"),
	}
}
