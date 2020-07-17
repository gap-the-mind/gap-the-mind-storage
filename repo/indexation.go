package repo

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

func createIndexer() {
	mapping := bleve.NewIndexMapping()

	fmt.Println((mapping))
}
