package service

import (
	"github.com/k0marov/newsbot/internal/core/domain"
	"log"
	"time"
)

type NewsService struct {
}

func NewNewsService() *NewsService {
	return &NewsService{}
}

func (n *NewsService) GetNews() <-chan domain.NewsEntry {
	log.Println("starting to publish news...")
	ch := make(chan domain.NewsEntry)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			ch <- domain.NewsEntry{URL: "Hello, test message"}
		}
	}()
	return ch
}
