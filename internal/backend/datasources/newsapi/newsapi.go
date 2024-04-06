package newsapi

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/k0marov/newsbot/internal/core/domain"
	"log"
	"net/url"
	"strconv"
)

type NewsAPI struct {
}

func NewNewsAPI() *NewsAPI {
	return &NewsAPI{}
}

func (n *NewsAPI) GetAllNews() ([]domain.NewsEntry, error) {
	// TODO: query other pages
	queryURL := constructSearchURL(1, pageSize, searchQuery).String()
	log.Println(queryURL)
	doc, err := htmlquery.LoadURL(queryURL)
	if err != nil {
		return nil, fmt.Errorf("loading news from api using htmlquery: %w", err)
	}
	nodes, err := htmlquery.QueryAll(doc, "/html/body/div[1]/div/div[2]/div/div[4]/div[2]/div/div[2]/form/div[*]/div/div/ul/li[3]")
	if err != nil {
		return nil, fmt.Errorf("querying for isbns in html using htmlquery: %w", err)
	}
	news := make([]domain.NewsEntry, len(nodes))
	for i := range nodes {
		news[i] = domain.NewsEntry{URL: nodes[i].FirstChild.Data}
	}
	return news, nil
}

func constructSearchURL(pageIndex, pageSize int, searchQuery string) *url.URL {
	u, err := url.Parse(searchURL)
	if err != nil {
		log.Panicf("unable to parse url: %v", err)
	}
	q := u.Query()
	q.Add("page", strconv.Itoa(pageIndex))
	q.Add("pageUnit", strconv.Itoa(pageSize))
	q.Add("schStr", searchQuery) // url.QueryEscape(searchQuery))
	u.RawQuery = q.Encode()
	return u
}
