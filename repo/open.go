package repo

import (
	"fmt"
	"github.com/gap-the-mind/gap-the-mind-storage/entity"
	"path/filepath"
	"time"

	"github.com/blevesearch/bleve/v2"
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
func Open(path string, providers []entity.Provider) (Storage, error) {
	defer logger.Sync()

	repo, err := git.PlainOpen(path)

	if err != nil {
		logger.Infow("No repo found - initialization",
			"entityPath",
			path,
		)

		repo, err = git.PlainInit(path, false)

		if err != nil {
			logger.Fatal("Failed to init repo",
				"entityPath",
				path,
			)
		}
	}

	logger.Infow("Found repo",
		"entityPath",
		path,
	)

	var index bleve.Index
	indexPath := filepath.Join(path, ".index")

	logger.Infow("Open index", "entityPath", indexPath)
	index, err = bleve.Open(indexPath)

	if err != nil {
		logger.Infow("No index found, creating a new one", "entityPath", indexPath)
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexPath, mapping)
	}

	commitChan := make(chan string, 100)

	storage := Storage{
		repo:            repo,
		indexer:         index,
		commits:         commitChan,
		entityProviders: providers,
	}

	return storage, err
}
