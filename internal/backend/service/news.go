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
	ds    NewsDS
	saver NewsSaver
}

func NewNewsService(ds NewsDS, saver NewsSaver) *NewsService {
	return &NewsService{ds, saver}
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
			newNews, err := n.saver.SaveAndReturnNew(news)
			if err != nil {
				log.Println("ERROR:", "failed saving fetched news and filtering only those that were not saved before", err)
			}
			log.Println("got", len(newNews), "new news")
			for _, newsEntry := range newNews {
				ch <- newsEntry
			}
		}
	}()
	return ch
}
