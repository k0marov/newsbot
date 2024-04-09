package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	BotToken           string
	BotCorrectPassword string
	NewsPollInterval   time.Duration
}

var config *Config

func GetConfig() *Config {
	var err error
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
	config.NewsPollInterval, err = time.ParseDuration(getEnvWithDefault("NEWS_POLL_INTERVAL", "5m"))
	if err != nil {
		log.Fatalf("failed parsing NEWS_POLL_INTERVAL as duration")
	}
	return config
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("please provide %s in env", key)
	}
	return value
}

func getEnvWithDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
