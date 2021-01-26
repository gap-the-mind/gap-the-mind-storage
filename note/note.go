package note

import (
	"github.com/gap-the-mind/gap-the-mind-storage/entity"
	"github.com/google/uuid"
)

type Note struct {
	id   string
	text string
}

func (n Note) Id() string {
	return n.id
}

func (n Note) SetId(s string) {
	n.id = s
}

func (n Note) Nature() string {
	return "note"
}

type NoteProvider struct {
}

// New creates a new note with the given text
func New(text string) *Note {
	id, _ := uuid.NewRandom()

	return &Note{
		id:   id.String(),
		text: text,
	}
}

func (provider NoteProvider) Accept(id string, nature string) entity.Entity {
	if nature != "note" {
		return nil
	}

	return Note{id: id}
}
