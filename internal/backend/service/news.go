package service

import (
	"github.com/k0marov/newsbot/internal/core/domain"
	"log"
	"time"
)

type NewsDS interface {
	GetAllNews() ([]domain.NewsEntry, error)
}

type NewsService struct {
	ds NewsDS
}

func NewNewsService(ds NewsDS) *NewsService {
	return &NewsService{ds}
}

func (n *NewsService) GetNews() <-chan domain.NewsEntry {
	log.Println("starting to publish news...")
	ch := make(chan domain.NewsEntry)
	go func() {
		for {
			time.Sleep(30 * time.Second)
			log.Println("getting all news from api...")
			news, err := n.ds.GetAllNews()
			if err != nil {
				log.Println("ERROR:", "failed getting all news from api:", err)
			}
			news = news[:5] // TODO: this is for a test
			for _, newsEntry := range news {
				ch <- newsEntry
			}
		}
	}()
	return ch
}
