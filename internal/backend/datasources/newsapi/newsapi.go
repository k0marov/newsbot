package newsapi

import "github.com/k0marov/newsbot/internal/core/domain"

type NewsAPI struct {
}

func NewNewsAPI() *NewsAPI {
	return &NewsAPI{}
}

func (n *NewsAPI) GetAllNews() ([]domain.NewsEntry, error) {
	return []domain.NewsEntry{{"Hello 1"}, {"Hello 2"}, {"Hello 3"}}, nil
}
