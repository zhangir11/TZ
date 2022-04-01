package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	ListenAddress string
}

type Config struct {
	Server                ServerConfig
	TestMode              bool
	SecretKey             string
	DatabaseURI           string
	DatabaseName          string
	AccessTokenTimeDelta  string
	RefreshTokenTimeDelta string
}

var Conf *Config

//-----------------------------------------------------------------------------------------------

func NewConfig() {
	godotenv.Load(".env")
	Conf = &Config{
		Server: ServerConfig{
			ListenAddress: getEnv("LISTEN_ADDRESS", "localhost:8080"),
		},
		SecretKey:             getEnv("SECRET_KEY", "very secret key"),
		DatabaseURI:           getEnv("DATABASE_URI", ""),
		AccessTokenTimeDelta:  getEnv("ACCESS_TOKEN_TIME_DELTA", "15m"),
		RefreshTokenTimeDelta: getEnv("REFRESH_TOKEN_TIME_DELTA", "24h"),
	}
	log.Println(Conf.DatabaseURI, getEnv("DATABASE_URI", ""), "hehe")
}

//-----------------------------------------------------------------------------------------------

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
