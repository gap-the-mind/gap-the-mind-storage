package repo

import (
	"github.com/gap-the-mind/gap-the-mind-storage/entity"
	"os"

	"github.com/blevesearch/bleve/v2"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
)

// StorageUnit wraps stored content
type storageUnit struct {
	ID      string
	Nature  string
	Content entity.Entity
}

// Layout is the layout
type Layout interface {
	Path(fs billy.Filesystem, entity entity.Entity)
	Id(file os.FileInfo)
}

// Storage is the storage
type Storage struct {
	repo            *git.Repository
	layout          Layout
	indexer         bleve.Index
	commits         chan string
	entityProviders []entity.Provider
}
