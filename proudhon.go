package main

import (
	"gitlab.com/ekai/proudhon/gap-the-mind-storage/log"
	"gitlab.com/ekai/proudhon/gap-the-mind-storage/repo"
	"gitlab.com/ekai/proudhon/gap-the-mind-storage/server"
)

func main() {
	logger := log.CreateLogger()
	defer logger.Sync()

	repoPath := "./storage/repo"

	logger.Infow("Open repo", "repo", repoPath)
	repo, err := repo.Open(repoPath)

	if err != nil {
		logger.Fatalw("Fail to oppen repo", "repo", repoPath)
	}

	server.StartServer(repo)
}
