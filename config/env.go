package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host      string
	Port      string
	DBUser    string
	DBPasswd  string
	DBName    string
	DBAddress string
}


var Envs = initConfig()
func initConfig() Config {
	godotenv.Load()
	return Config{
		Host:      getEnv("DB_HOST", "localhost"),
        Port:      getEnv("DB_PORT", "3306"),
        DBUser:    getEnv("DB_USER", "root"),
        DBPasswd:  getEnv("DB_PASSWD", "root"),
        DBName:    getEnv("DB_NAME", "go_ecom"),
        DBAddress: getEnv("DB_ADDRESS", "127.0.0.1:3306"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}