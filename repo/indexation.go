package repo

import (
	"github.com/blevesearch/bleve/v2"
	"os"
	"path"
	"path/filepath"
)

// Reindex re-indexes
func (s Storage) Reindex() error {
	fs, err := s.fs()

	if err != nil {
		return err
	}

	err = filepath.Walk(fs.Root(), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && (info.Name() == ".index" || info.Name() == ".git") {
			return filepath.SkipDir
		}

		ext := path.Ext(info.Name())

		logger.Debugw("Reindex candidate", "file", info.Name(), "extension", ext)

		if info.IsDir() || ext != ".toml" {
			return nil
		}

		relPath, err := filepath.Rel(fs.Root(), p)

		if err != nil {
			return err
		}

		logger.Infow("Reindex", "file", relPath)
		entity, err := s.Read(relPath)

		logger.Debugw("Reindex ", "entity", entity)

		if err != nil {
			return err
		}

		return s.indexer.Index((*entity).Id(), entity)
	})

	return err

}

func (s *Storage) Search(search *bleve.SearchRequest) error {
	results, err := s.indexer.Search(search)

	if err != nil {
		return err
	}

	logger.Infow("Search", "results", results)

	return nil
}
