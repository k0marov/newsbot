package main

import (
	"github.com/k0marov/newsbot/internal/backend/datasources/newsapi"
	"github.com/k0marov/newsbot/internal/backend/datasources/repository"
	"github.com/k0marov/newsbot/internal/backend/service"
	"github.com/k0marov/newsbot/internal/core/config"
	"github.com/k0marov/newsbot/internal/frontend"
)

func main() {
	cfg := config.GetConfig()

	newsDS := newsapi.NewNewsAPI()
	newsSVC := service.NewNewsService(newsDS)
	newsCh := newsSVC.GetNews()

	repo := repository.NewRepository()
	authSVC := service.NewAuthService(cfg.BotCorrectPassword, repo)

	frontend.StartBot(cfg.BotToken, newsCh, authSVC, authSVC)
}
