package repo

import (
	"log"

	"github.com/go-git/go-git/v5"
)

func create(path string) *git.Repository {
	repo, err := git.PlainInit(path, false)

	if err != nil {
		log.Fatal(err)
	}

	return repo
}
