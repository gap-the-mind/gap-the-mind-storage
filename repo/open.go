package repo

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/go-git/go-git/v5"
)

func commit(commits <-chan string, tree *git.Worktree, debounce int64) error {

	var msg string
	var start int64

	for c := range commits {
		msg = c
		start = time.Now().Unix()

		for time.Now().Unix()-start < debounce {
			select {
			case c = <-commits:
				msg += "\n" + c
			default:
				time.Sleep(1 * time.Second)
			}
		}

		_, err := tree.Add(".")

		if err != nil {
			return err
		}

		_, err = tree.Commit(msg, &git.CommitOptions{})

		if err != nil {
			return fmt.Errorf("Failed to commit: %w", err)
		}
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
