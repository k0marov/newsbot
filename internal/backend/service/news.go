package service

import (
	"log"
	"time"
)

type NewsService struct {
}

func NewNewsService() *NewsService {
	return &NewsService{}
}

func (n *NewsService) GetNews() <-chan string {
	log.Println("starting to publish news...")
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			ch <- "Hello, test message"
		}
	}()
	return ch
}
