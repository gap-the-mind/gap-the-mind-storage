package server

import (
	"github.com/go-git/go-git/v5"
	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.com/ekai/proudhon/gap-the-mind-storage/log"
)

// StartServer starts the server
func StartServer(repo *git.Repository) {
	logger := log.CreateLogger()
	defer logger.Sync()

	logger.Infow("Start server",
		"schema", &graphql.Schema{})
}
