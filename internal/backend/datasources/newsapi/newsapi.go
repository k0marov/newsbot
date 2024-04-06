package newsapi

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/k0marov/newsbot/internal/core/domain"
	"log"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

type NewsAPI struct {
}

func NewNewsAPI() *NewsAPI {
	return &NewsAPI{}
}

func (n *NewsAPI) GetAllNews() ([]domain.NewsEntry, error) {
	news := make([]domain.NewsEntry, 0, 1000)
	for pageIndex := 1; pageIndex <= 3; pageIndex++ {
		fetchedNews, err := fetchNews(pageIndex)
		if err != nil {
			return nil, fmt.Errorf("fetching news for page #%d: %w", pageIndex, err)
		}
		log.Printf("got %d news for page %d\n", len(fetchedNews), pageIndex)
		if len(fetchedNews) == 0 {
			break
		}
		news = append(news, fetchedNews...)
	}
	log.Printf("got a total of %d news\n", len(news))
	return news, nil
}

func fetchNews(pageIndex int) ([]domain.NewsEntry, error) {
	queryURL := constructSearchURL(pageIndex, pageSize, searchQuery).String()
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
		isbnText := nodes[i].FirstChild.Data
		isbn := parseISBN(isbnText)
		news[i] = domain.NewsEntry{URL: constructNewsEntryURL(isbn)}
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

// parseISBN parses extracts ISBN number from a string that looks like "ISBN: 979-11-93265-28-4 (75810)"
// NOTE: of course, it will break in the future, when website's design changes
func parseISBN(isbnText string) string {
	const isbnLength = 13

	braceInd := strings.Index(isbnText, "(")
	if braceInd == -1 {
		log.Panicf("failed parsing isbn %q: brace not found", isbnText)
	}
	isbnText = isbnText[:braceInd]
	isbnRunes := make([]rune, 0, isbnLength)
	for _, c := range isbnText {
		if unicode.IsDigit(c) {
			isbnRunes = append(isbnRunes, c)
		}
	}
	if len(isbnRunes) != isbnLength {
		log.Panicf("failed parsing isbn %q: got invalid length", isbnText)
	}
	return string(isbnRunes)
}
