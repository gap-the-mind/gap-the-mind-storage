package repo

import (
	"fmt"
	"time"

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
