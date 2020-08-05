package repo

import (
	"github.com/blevesearch/bleve"
	"github.com/go-git/go-git/v5"
)

// StorageUnit wraps stored content
type storageUnit struct {
	ID      string
	Nature  string
	Content EntityRef
}

// EntityRef is base for all storable objects
type EntityRef interface {
	Id() string
	SetId(string)

	Nature() string
}

// Storage is the storage
type Storage struct {
	Layout

	repo    *git.Repository
	indexer *bleve.Index
	commits chan string
}
