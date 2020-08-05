package repo

import (
	"github.com/blevesearch/bleve"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

// OpenMemory open or create a repo in memory
func OpenMemory() (Storage, error) {
	storage := memory.NewStorage()

	repo, err := git.Open(storage, nil)

	logger.Infow("No repo found - initialization in memory")

	repo, err = git.Init(storage, memfs.New())

	if err != nil {
		logger.Errorw("Failed to init in memory repo")

		return Storage{}, err
	}

	logger.Infow("Found repo")

	return createStorage(repo)

}

// OpenFilesystem open or create a repo
func OpenFilesystem(path string) (Storage, error) {
	defer logger.Sync()

	repo, err := git.PlainOpen(path)

	if err != nil {
		logger.Infow("No repo found - initialization",
			"path",
			path,
		)

		repo, err = git.PlainInit(path, false)

		if err != nil {
			logger.Errorw("Failed to init repo",
				"path",
				path,
			)

			return Storage{}, err
		}
	}

	logger.Infow("Found repo",
		"path",
		path,
	)

	return createStorage(repo)
}

// Open a new repo
func createStorage(repo *git.Repository) (Storage, error) {

	var index bleve.Index
	tree, err := repo.Worktree()

	if err != nil {
		return Storage{}, err
	}

	fs := tree.Filesystem

	indexPath := fs.Join(fs.Root(), ".index")

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
