package domain

import "time"

type NewsEntry struct {
	URL             string
	PublicationDate time.Time
}
