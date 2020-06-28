package repo

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pelletier/go-toml"
)

var layout = map[string]string{
	"note": "notes",
}

func prefix(fs billy.Filesystem, typ string) string {
	prefix, found := layout[typ]

	if !found {
		prefix = "miscellanous"
	}

	return prefix
}

func filename(typ string, id string) string {
	return fmt.Sprintf("%s.toml", id)
}

func id(typ string, filename string) string {
	return strings.TrimSuffix(filename, ".toml")
}

func path(fs billy.Filesystem, typ string, id string) string {
	path := fs.Join(prefix(fs, typ), filename(typ, id))

	return path

}

func (s *Storage) Create(typ string, id string, content interface{}) error {
	tree, err := s.repo.Worktree()

	if err != nil {
		return err
	}

	fs := tree.Filesystem

	b, err := toml.Marshal(content)

	if err != nil {
		return err
	}

	path := path(fs, typ, id)
	file, err := fs.Create(path)

	if err != nil {
		return err
	}

	_, err = file.Write(b)

	if err != nil {
		return err
	}

	_, err = tree.Add(path)

	if err != nil {
		return err
	}

	_, err = tree.Commit(fmt.Sprintf("Create note %s at %s", id, path), &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		}})

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Get(typ string, id string, target interface{}) error {
	tree, err := s.repo.Worktree()
	fs := tree.Filesystem

	if err != nil {
		return fmt.Errorf("Filesystem error: %w", err)
	}

	path := path(fs, typ, id)

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
		return fmt.Errorf("Failed to unmarshal %s from %s: %w", id, path, err)
	}

	return nil

}

func (s *Storage) Update(typ string, id string, content interface{}) error {
	logger.Debugw("Update",
		"type", typ,
		"id", id,
		"content", content,
	)

	tree, err := s.repo.Worktree()
	fs := tree.Filesystem

	if err != nil {
		return fmt.Errorf("Filesystem error: %w", err)
	}

	b, err := toml.Marshal(content)

	if err != nil {
		return fmt.Errorf("Failed to marshal: %w", err)
	}

	path := path(fs, typ, id)

	logger.Debugw("Saving",
		"type", typ,
		"id", id,
		"path", path,
		"content", b,
	)

	file, err := fs.Create(path)

	if err != nil {
		return fmt.Errorf("Failed to open file %s: %w", path, err)
	}

	_, err = file.Write(b)

	if err != nil {
		return fmt.Errorf("Failed to write file %s: %w", path, err)
	}

	_, err = tree.Add(path)

	if err != nil {
		return fmt.Errorf("Failed to add file %s: %w", path, err)
	}

	_, err = tree.Commit(fmt.Sprintf("Update note %s at %s", id, path), &git.CommitOptions{
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

func (s *Storage) List(typ string) ([]string, error) {
	tree, err := s.repo.Worktree()
	fs := tree.Filesystem

	if err != nil {
		return nil, err
	}

	infos, err := fs.ReadDir(prefix(fs, typ))

	if err != nil {
		return nil, err
	}

	ids := make([]string, len(infos))

	for i, info := range infos {
		ids[i] = id(typ, info.Name())
	}

	return ids, nil
}
