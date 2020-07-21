package repo

import (
	"fmt"
	"io/ioutil"

	"github.com/go-git/go-billy/v5"
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

func readEntity(file billy.File, target EntityRef) error {
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

func writeEntity(file billy.File, content EntityRef) error {
	unit := storageUnit{
		ID:      content.Id(),
		Nature:  content.Nature(),
		Content: content,
	}

	b, err := toml.Marshal(unit)

	if err != nil {
		return err
	}

	_, err = file.Write(b)

	return err
}

// Delete deletes an entity
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

// Save creates and entity
func (s *Storage) Save(content EntityRef) error {
	content.SetId(uuid.New().String())

	fs, err := s.fs()

	if err != nil {
		return err
	}

	path := path(fs, content)
	file, err := fs.Create(path)

	err = writeEntity(file, content)

	return err
}

// Get retreive an entity
func (s *Storage) Get(target EntityRef) error {
	fs, err := s.fs()

	if err != nil {
		return fmt.Errorf("Filesystem error: %w", err)
	}

	path := path(fs, target)
	file, err := fs.Open(path)

	err = readEntity(file, target)

	return err

}

// List lists all entities of the given type
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
