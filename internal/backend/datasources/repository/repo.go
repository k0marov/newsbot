package repository

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

type Repository struct {
	db *leveldb.DB
}

func NewRepository() *Repository {
	db, err := leveldb.OpenFile("leveldb/storage", nil)
	if err != nil {
		log.Fatalf("opening leveldb: %v", err)
	}
	return &Repository{db}
}
