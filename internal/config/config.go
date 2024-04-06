package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	BotToken           string
	BotCorrectPassword string
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}
	if err := godotenv.Load(); err != nil {
		// it's ok if there is no dotenv, but let's warn
		log.Printf("could not load dotenv: %v", err)
	}
	config = &Config{}
	config.BotToken = mustGetEnv("BOT_TOKEN")
	config.BotCorrectPassword = mustGetEnv("BOT_CORRECT_PASSWORD")
	return config
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("please provide %s in env", key)
	}
	return value
}
