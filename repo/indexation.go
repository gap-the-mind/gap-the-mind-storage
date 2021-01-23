package repo

import (
	"os"
	"path/filepath"
)

// Reindex reindexes
func (s Storage) Reindex() error {
	logger.Infow("Reindex", "path", path)
	fs, err := s.fs()

	if err != nil {
		return err
	}

	err = filepath.Walk(fs.Root(), func(path string, info os.FileInfo, err error) error {

		if info.IsDir() && (info.Name() == ".index" || info.Name() == ".git") {
			return filepath.SkipDir
		}

		if err != nil {
			return err
		}

		logger.Info(path, info.Size())
		return nil
	})

	return err

}
