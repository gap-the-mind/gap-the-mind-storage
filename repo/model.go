package repo

import (
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
)

// EntityRef is base for all storable objects
type EntityRef interface {
	Id() string
	SetId(string)

	Nature() string
}

type Layout interface {
	Path(fs billy.Filesystem, entity EntityRef)
	Id(file os.FileInfo)
}

type Storage struct {
	repo *git.Repository

	layout Layout
}
