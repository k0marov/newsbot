package main

import (
	"github.com/k0marov/newsbot/internal/backend/service"
	"github.com/k0marov/newsbot/internal/config"
	"github.com/k0marov/newsbot/internal/frontend"
)

func main() {
	cfg := config.GetConfig()
	newsSVC := service.NewNewsService()
	newsCh := newsSVC.GetNews()
	authSVC := service.NewAuthService(cfg.BotCorrectPassword)
	frontend.StartBot(cfg.BotToken, newsCh, authSVC, authSVC)
}
