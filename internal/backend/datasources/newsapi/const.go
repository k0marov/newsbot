package newsapi

import "fmt"

const (
	searchURL   = "https://www.nl.go.kr/seoji/contents/S80100000000.do?schType=simple&ebookYn=Y#link"
	searchQuery = "삽화가(그림작가)"
	pageSize    = 10
)

func constructNewsEntryURL(isbn string) string {
	return fmt.Sprintf("https://www.nl.go.kr/seoji/contents/S80100000000.do?schM=intgr_detail_view_isbn&isbn=%s", isbn)
}
