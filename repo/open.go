package repo

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func (s *Storage) commit(commits <-chan string, tree *git.Worktree) error {

	var msg string

	for c := range commits {
		msg
	}
	_, err := tree.Add(".")

	if err != nil {
		return err
	}

	_, err = tree.Commit(msg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		}})

	if err != nil {
		return fmt.Errorf("Failed to commit: %w", err)
	}

	return nil
}

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

	var index bleve.Index
	indexPath := filepath.Join(path, ".index")

	logger.Infow("Open index", "path", indexPath)

	if err != nil {
		logger.Info("Create new index")

		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexPath, mapping)
	}

	commitChan := make(chan string, 100)

	storage := Storage{repo: repo, indexer: &index, commits: commitChan}

	return storage, err
}
