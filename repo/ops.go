package repo

import (
	"fmt"
	"github.com/gap-the-mind/gap-the-mind-storage/entity"
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

func readEntity(file billy.File, target entity.Entity) error {
	b, err := ioutil.ReadAll(file)

	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", file.Name(), err)
	}

	err = toml.Unmarshal(b, target)

	if err != nil {
		return fmt.Errorf("failed to unmarshal %s from %s: %w", target.Id(), file.Name(), err)
	}

	return nil
}

func writeEntity(file billy.File, content entity.Entity) error {
	b, err := toml.Marshal(content)

	if err != nil {
		return err
	}

	_, err = file.Write(b)

	return err
}

func (s *Storage) commit(message string) {
	s.commits <- message
}

// Delete deletes an entity
func (s *Storage) Delete(target entity.Entity) error {
	fs, err := s.fs()

	if err != nil {
		return err
	}

	path := entityPath(fs, target)

	err = fs.Remove(path)

	if err != nil {
		return err
	}

	s.commit(fmt.Sprintf("Delete note %s at %s", target.Id(), path))

	return nil
}

// Save creates and entity
func (s *Storage) Save(content entity.Entity) error {
	content.SetId(uuid.New().String())

	fs, err := s.fs()

	if err != nil {
		return err
	}

	path := entityPath(fs, content)
	file, err := fs.Create(path)

	err = writeEntity(file, content)

	if err != nil {
		return err
	}

	s.commit(fmt.Sprintf("Save note %s at %s", content.Id(), path))

	return nil
}

// Get retreive an entity
func (s *Storage) Get(target entity.Entity) error {
	fs, err := s.fs()

	if err != nil {
		return fmt.Errorf("filesystem error: %w", err)
	}

	path := entityPath(fs, target)
	file, err := fs.Open(path)

	err = readEntity(file, target)

	return err

}

// Get retreive an entity
func (s *Storage) Read(path string) (*entity.Entity, error) {
	fs, err := s.fs()

	if err != nil {
		return nil, fmt.Errorf("filesystem error: %w", err)
	}

	file, err := fs.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", file.Name(), err)
	}

	tree, err := toml.LoadBytes(b)

	if err != nil {
		return nil, fmt.Errorf("parsing error: %w", err)
	}

	id := tree.Get("Id")
	nature := tree.Get("Nature")
	var target entity.Entity

	for _, p := range s.entityProviders {
		target = p.Accept(id.(string), nature.(string))

		if target != nil {
			err = tree.Unmarshal(&target)

			if err != nil {
				return &target, nil
			}

		}
	}

	return nil, nil

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
