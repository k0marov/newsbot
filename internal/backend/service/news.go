package service

import (
	"github.com/k0marov/newsbot/internal/core/domain"
	"log"
	"time"
)

type NewsDS interface {
	GetAllNews() ([]domain.NewsEntry, error)
}

type NewsSaver interface {
	SaveAndReturnNew(fetchedNews []domain.NewsEntry) ([]domain.NewsEntry, error)
}

type NewsService struct {
	ds           NewsDS
	saver        NewsSaver
	pollInterval time.Duration
}

func NewNewsService(ds NewsDS, saver NewsSaver, pollInterval time.Duration) *NewsService {
	return &NewsService{ds, saver, pollInterval}
}

func (n *NewsService) GetNews() <-chan domain.NewsEntry {
	log.Println("starting to publish news...")
	ch := make(chan domain.NewsEntry)
	go func() {
		for {
			time.Sleep(n.pollInterval)
			log.Println("getting all news from api...")
			news, err := n.ds.GetAllNews()
			if err != nil {
				log.Println("ERROR:", "failed getting all news from api:", err)
			}
			newNews, err := n.saver.SaveAndReturnNew(news)
			if err != nil {
				log.Println("ERROR:", "failed saving fetched news and filtering only those that were not saved before", err)
			}
			log.Println("got", len(newNews), "new news")
			for _, newsEntry := range news {
				ch <- newsEntry
				time.Sleep(time.Second)
			}
		}
	}()
	return ch
}
