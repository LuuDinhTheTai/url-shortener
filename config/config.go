package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	URI string
}

func LoadEnv() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	uri := os.Getenv("DATABASE_URI")

	return Config{
		Server: ServerConfig{
			Host: host,
			Port: port,
		},
		Database: DatabaseConfig{
			URI: uri,
		},
	}
}
