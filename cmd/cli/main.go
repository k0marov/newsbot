package main

import (
	"github.com/k0marov/newsbot/internal/backend"
	"github.com/k0marov/newsbot/internal/config"
	"github.com/k0marov/newsbot/internal/frontend"
	"log"
)

func main() {
	cfg := config.GetConfig()
	svc := backend.NewAuthService(cfg.BotCorrectPassword)
	log.Println("Starting bot...")
	frontend.StartBot(cfg.BotToken, svc)
}
