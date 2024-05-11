package newsapi

import "fmt"

const (
	searchURL     = "https://www.nl.go.kr/seoji/contents/S80100000000.do?schType=simple"
	refererHeader = "https://nl.go.kr/seoji/contents/S80100000000.do?schType=simple&schStr=%EC%82%BD%ED%99%94%EA%B0%80%28%EA%B7%B8%EB%A6%BC%EC%9E%91%EA%B0%80%29+%3A"
	searchQuery   = "삽화가(그림작가) :"
	pageSize      = 100
)

func constructNewsEntryURL(isbn string) string {
	return fmt.Sprintf("https://www.nl.go.kr/seoji/contents/S80100000000.do?schM=intgr_detail_view_isbn&isbn=%s", isbn)
}
