package main

import (
	"github.com/k0marov/newsbot/internal/backend/datasources/newsapi"
	"github.com/k0marov/newsbot/internal/backend/datasources/repository"
	"github.com/k0marov/newsbot/internal/backend/service"
	"github.com/k0marov/newsbot/internal/core/config"
	"github.com/k0marov/newsbot/internal/frontend"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()

	newsDS := newsapi.NewNewsAPI()
	newsSaver := repository.NewNewsSaver()

	newsSVC := service.NewNewsService(newsDS, newsSaver, cfg.NewsPollInterval)
	newsCh := newsSVC.GetNews()

	repo, closeRepo := repository.NewRepository()
	authSVC := service.NewAuthService(cfg.BotCorrectPassword, repo)

	done := make(chan struct{}, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	stopBot := frontend.StartBot(cfg.BotToken, newsCh, authSVC, authSVC)

	go func() {
		<-sigs

		log.Println("Stopping bot...")
		stopBot()
		log.Println("Closing repo...")
		if err := closeRepo(); err != nil {
			log.Println("ERROR:", "failed closing repo:", err.Error())
		}

		done <- struct{}{}
	}()

	log.Println("handling bot updates...")

	<-done
	log.Println("Done")

}
