package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBNAME     string
}

type APIConfig struct {
	ServerAddress string
}

type Config struct {
	DBConfig
	APIConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbport, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return fmt.Errorf("Invalid port %s %v", dbport, err)
	}

	c.DBConfig = DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     dbport,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBNAME:     os.Getenv("DB_NAME"),
	}

	c.APIConfig = APIConfig{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}

	if c.DBHost == "" || c.DBPort == 0 || c.DBUser == "" || c.DBPassword == "" || c.DBNAME == "" || c.ServerAddress == "" {
		return fmt.Errorf("required config")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}
