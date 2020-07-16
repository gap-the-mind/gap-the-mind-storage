package repo

import (
	"github.com/go-git/go-git/v5"
)

// Storage provide CRUD ops

// Open a new repo
func Open(path string) (Storage, error) {
	defer logger.Sync()

	repo, err := git.PlainOpen(path)

	if err != nil {
		logger.Infow("No repo found - initialization",
			"path",
			path,
		)

		repo, err = git.PlainInit(path, false)

		if err != nil {
			logger.Fatal("Failed to init repo",
				"path",
				path,
			)
		}
	}

	logger.Infow("Found repo",
		"path",
		path,
	)

	return Storage{repo: repo}, nil
}
