package repo

import (
	"github.com/go-git/go-git/v5"
	"go.uber.org/zap"
)

// Open a new repo
func Open(path string) (*git.Repository, error) {
	l, _ := zap.NewDevelopment()
	logger := l.Sugar()

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

	return repo, nil
}
