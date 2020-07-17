package repo

import (
	"os"

	"github.com/blevesearch/bleve"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
)

// EntityRef is base for all storable objects
type EntityRef interface {
	Id() string
	SetId(string)

	Nature() string
}

// Layout is the layout
type Layout interface {
	Path(fs billy.Filesystem, entity EntityRef)
	Id(file os.FileInfo)
}

// Storage is the storage
type Storage struct {
	repo    *git.Repository
	layout  Layout
	indexer *bleve.Index
}
