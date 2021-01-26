package main

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/gap-the-mind/gap-the-mind-storage/entity"
	log "github.com/gap-the-mind/gap-the-mind-storage/log"
	"github.com/gap-the-mind/gap-the-mind-storage/note"

	"github.com/gap-the-mind/gap-the-mind-storage/repo"
)

var logger = log.CreateLogger()

func main() {

	repoPath := "../storage"

	logger.Infow("Open repo", "path", repoPath)
	repository, err := repo.Open(
		repoPath,
		[]entity.Provider{
			note.NoteProvider{},
		})

	if err != nil {
		logger.Panic(err)
	}

	n := note.New("Toto")

	repository.Save(n)

	// err = repository.Reindex()

	query := bleve.NewQueryStringQuery("Toto")
	search := bleve.NewSearchRequestOptions(query, 10, 0, true)
	repository.Search(search)

}
