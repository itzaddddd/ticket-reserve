package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Postgres DbConfig
}

type AppConfig struct {
	Host string
	Port string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("load env failed cuz: %v", err)
		panic(err)
	}

	return &Config{
		App: AppConfig{
			Port: os.Getenv("APP_PORT"),
		},
		Postgres: DbConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DbName:   os.Getenv("POSTGRES_DB"),
		},
	}
}
