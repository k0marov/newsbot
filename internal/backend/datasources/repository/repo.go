package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"log"
	"slices"
)

const authenticatedKey = "authenticatedChatIDs"

type Repository struct {
	db *leveldb.DB
}

func NewRepository() (r *Repository, close func() error) {
	db, err := leveldb.OpenFile("/leveldb", nil)
	if err != nil {
		log.Fatalf("opening leveldb: %v", err)
	}
	return &Repository{db}, db.Close
}

// NOTE: this is highly inefficient. If this becomes a burden, migrate to PostgreSQL

func (r *Repository) AddToAuthenticated(chatID string) error {
	authenticated, err := r.GetAllAuthenticated()
	if err != nil {
		return fmt.Errorf("failed getting current authenticated list: %w", err)
	}
	if slices.Contains(authenticated, chatID) {
		log.Printf("didn't add %q to list of authenticated, because it is already there", chatID)
		return nil
	}
	authenticated = append(authenticated, chatID)
	authenticatedJSON, err := json.Marshal(authenticated)
	if err != nil {
		return fmt.Errorf("failed marshalling json: %w", err)
	}
	if err := r.db.Put([]byte(authenticatedKey), authenticatedJSON, &opt.WriteOptions{}); err != nil {
		return fmt.Errorf("failed writing to leveldb: %w", err)
	}
	return nil
}

func (r *Repository) GetAllAuthenticated() ([]string, error) {
	res, err := r.db.Get([]byte(authenticatedKey), &opt.ReadOptions{})
	if errors.Is(err, leveldb.ErrNotFound) {
		return make([]string, 0), nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed GETting in leveldb: %w", err)
	}
	var chatIDs []string
	if err := json.Unmarshal(res, &chatIDs); err != nil {
		return nil, fmt.Errorf("failed unmarshalling json: %w", err)
	}
	return chatIDs, nil
}
