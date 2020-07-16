package repo

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/uuid"
	"github.com/pelletier/go-toml"
)

func (s *Storage) fs() (billy.Filesystem, error) {
	tree, err := s.repo.Worktree()

	if err != nil {
		return nil, err
	}

	return tree.Filesystem, nil
}

func (s *Storage) commit(path string, msg string) error {
	tree, err := s.repo.Worktree()
	_, err = tree.Add(path)

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

func (s *Storage) Delete(target EntityRef) error {
	fs, err := s.fs()

	if err != nil {
		return err
	}

	path := path(fs, target)

	err = fs.Remove(path)

	if err != nil {
		return err
	}

	return s.commit(path, fmt.Sprintf("Delete note %s at %s", target.Id(), path))
}

func (s *Storage) Create(content EntityRef) error {
	content.SetId(uuid.New().String())

	fs, err := s.fs()

	if err != nil {
		return err
	}

	b, err := toml.Marshal(content)

	if err != nil {
		return err
	}

	path := path(fs, content)
	file, err := fs.Create(path)

	if err != nil {
		return err
	}

	_, err = file.Write(b)

	if err != nil {
		return err
	}

	return s.commit(path, fmt.Sprintf("Create note %s at %s", content.Id(), path))
}

func (s *Storage) Get(target EntityRef) error {
	fs, err := s.fs()

	if err != nil {
		return fmt.Errorf("Filesystem error: %w", err)
	}

	path := path(fs, target)

	file, err := fs.Open(path)

	if err != nil {
		return fmt.Errorf("Failed to open file %s: %w", path, err)
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return fmt.Errorf("Failed to read file %s: %w", path, err)
	}

	err = toml.Unmarshal(b, target)

	if err != nil {
		return fmt.Errorf("Failed to unmarshal %s from %s: %w", target.Id(), path, err)
	}

	return nil

}

func (s *Storage) Update(content EntityRef) error {
	logger.Debugw("Update",
		"type", content.Nature(),
		"id", content.Id(),
	)

	fs, err := s.fs()

	if err != nil {
		return fmt.Errorf("Filesystem error: %w", err)
	}

	b, err := toml.Marshal(content)

	if err != nil {
		return fmt.Errorf("Failed to marshal: %w", err)
	}

	path := path(fs, content)

	logger.Debugw("Saving",
		"type", content.Nature(),
		"id", content.Id(),
		"path", path,
	)

	file, err := fs.Create(path)

	if err != nil {
		return fmt.Errorf("Failed to open file %s: %w", path, err)
	}

	_, err = file.Write(b)

	if err != nil {
		return fmt.Errorf("Failed to write file %s: %w", path, err)
	}

	return s.commit(path, fmt.Sprintf("Update note %s at %s", content.Id(), path))

}

func (s *Storage) List(typ string) []string {
	tree, err := s.repo.Worktree()
	fs := tree.Filesystem

	if err != nil {
		return make([]string, 0)
	}

	infos, err := fs.ReadDir(prefix(typ))

	if err != nil {
		return make([]string, 0)
	}

	ids := make([]string, len(infos))

	for i, info := range infos {
		ids[i] = id(info.Name())
	}

	return ids
}
