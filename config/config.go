package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func getEnv(envName string, defaultValue string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	val, present := os.LookupEnv(envName)
	if !present {
		return defaultValue
	}
	return val
}

func GetConfig() *Config {
	Dialect := getEnv("DATABASE_CONNECTION", "postgres")
	Host := getEnv("DATABASE_HOST", "localhost")
	Port, _ := strconv.Atoi(getEnv("DATABASE_PORT", "5432"))
	Username := getEnv("DATABASE_USERNAME", "postgres")
	Password := getEnv("DATABASE_PASSWORD", "")
	Name := getEnv("DATABASE_NAME", "todopp")
	Charset := getEnv("DATABASE_CHARSET", "utf8")

	return &Config{
		DB: &DBConfig{
			Dialect:  Dialect,
			Host:     Host,
			Port:     Port,
			Username: Username,
			Password: Password,
			Name:     Name,
			Charset:  Charset,
		},
	}
}
