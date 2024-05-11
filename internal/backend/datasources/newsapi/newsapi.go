package newsapi

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/k0marov/newsbot/internal/core/domain"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type NewsAPI struct {
}

func NewNewsAPI() *NewsAPI {
	return &NewsAPI{}
}

func (n *NewsAPI) GetAllNews() ([]domain.NewsEntry, error) {
	var news []domain.NewsEntry
	pageIndex := 1
	for {
		log.Printf("fetching page %d\n", pageIndex)
		fetchedNews, err := fetchNews(pageIndex)
		if err != nil {
			return nil, fmt.Errorf("fetching news from website: %w", err)
		}
		if len(fetchedNews) == 0 {
			log.Printf("got zero news for page %d, stopping\n", pageIndex)
			break
		}
		news = append(news, fetchedNews...)
		pageIndex++
		oldestPostDate := oldestPost(fetchedNews).PublicationDate
		log.Printf("oldest post %v\n", oldestPostDate)
		if oldestPostDate.Before(time.Now()) {
			break
		}
	}
	log.Printf("got a total of %d news\n", len(news))
	return news, nil
}

func fetchNews(pageIndex int) ([]domain.NewsEntry, error) {
	resp, err := performSearchReq(pageIndex)
	if err != nil {
		return nil, fmt.Errorf("performing search request: %w", err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing response as html: %w", err)
	}
	nodes, err := htmlquery.QueryAll(doc, "/html/body/div[1]/div/div[2]/div/div[4]/div[2]/div/div[2]/form/div[*]/div/div/ul") // NOTE: of course, this will break someday
	if err != nil {
		return nil, fmt.Errorf("querying for isbns in html using htmlquery: %w", err)
	}
	news := make([]domain.NewsEntry, len(nodes))
	for i := range nodes { // 3 and 5
		isbnNode := htmlquery.FindOne(nodes[i], "/li[3]")
		isbn := parseISBN(isbnNode.FirstChild.Data)
		publicationDateNode := htmlquery.FindOne(nodes[i], "/li[last()-1]")
		publicationDate := parsePublicationDate(publicationDateNode.FirstChild.Data)
		price := htmlquery.FindOne(nodes[i], "/li[last()]")
		news[i] = domain.NewsEntry{
			URL:             constructNewsEntryURL(isbn),
			PublicationDate: publicationDate,
			Price:           price.FirstChild.Data,
		}
		log.Printf("%s\n", publicationDate)
	}
	return news, nil
}

func oldestPost(news []domain.NewsEntry) domain.NewsEntry {
	return slices.MaxFunc(news, func(a, b domain.NewsEntry) int { return a.PublicationDate.Compare(b.PublicationDate) })
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

func performSearchReq(pageIndex int) (*http.Response, error) {
	queryURL := constructSearchURL(pageIndex, pageSize, searchQuery).String()
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, queryURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed constructing request: %w", err)
	}
	req.Header.Add("Referer", refererHeader)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed executing request to website: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed executing request to website: got status %d", resp.StatusCode)
	}
	return resp, nil
}

// parseISBN extracts publication date from string that looks like this: "발행(예정)일: 2025.10.09"
// NOTE: of course, it will break in the future, when website's design changes
func parsePublicationDate(dateText string) time.Time {
	//log.Println(dateText)
	dateRunes := make([]rune, 0)
	for _, c := range dateText {
		if unicode.IsDigit(c) || c == '.' {
			dateRunes = append(dateRunes, c)
		}
	}
	date, err := time.Parse("2006.01.02", string(dateRunes))
	if err != nil {
		//log.Println("ERROR:", "failed parsing date", dateText, err)
		return time.Time{}
	}
	return date
}

// parseISBN extracts ISBN number from a string that looks like "ISBN: 979-11-93265-28-4 (75810)"
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
