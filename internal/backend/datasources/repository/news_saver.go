package repository

import (
	"github.com/k0marov/newsbot/internal/core/domain"
	"sync"
)

type NewsSaver struct {
	viewedNews   map[domain.NewsEntry]struct{}
	viewedNewsMu sync.Mutex
}

func NewNewsSaver() *NewsSaver {
	return &NewsSaver{
		viewedNews: make(map[domain.NewsEntry]struct{}),
	}
}

func (n *NewsSaver) SaveAndReturnNew(fetchedNews []domain.NewsEntry) (news []domain.NewsEntry, err error) {
	n.viewedNewsMu.Lock()
	defer n.viewedNewsMu.Unlock()

	if len(n.viewedNews) == 0 { // if we haven't saved any news yet, then this is the first chunk and it is not considered new
		n.markAsViewed(fetchedNews)
		return nil, nil
	}

	news = make([]domain.NewsEntry, 0, len(fetchedNews))
	for _, newsEntry := range fetchedNews {
		if _, viewed := n.viewedNews[newsEntry]; !viewed {
			news = append(news, newsEntry)
			n.viewedNews[newsEntry] = struct{}{}
		}
	}
	return news, nil
}

// markAsViewed is not thread-safe, it is an internal helper
func (n *NewsSaver) markAsViewed(news []domain.NewsEntry) {
	for _, newsEntry := range news {
		n.viewedNews[newsEntry] = struct{}{}
	}
}
