//go:generate easyjson

package note

import (
	"github.com/gap-the-mind/gap-the-mind-storage/entity"
	"github.com/google/uuid"
)

const noteNature = "note"

//easyjson:json
type Note struct {
	NoteId     string `json:"id"`
	Text       string `json:"text"`
	NoteNature string `json:"nature"`
}

func (n Note) Id() string {
	return n.NoteId
}

func (n Note) SetId(s string) {
	n.NoteId = s
}

func (n Note) Nature() string {
	return noteNature
}

type NoteProvider struct {
}

// New creates a new note with the given text
func New(text string) *Note {
	id, _ := uuid.NewRandom()

	return &Note{
		NoteId:     id.String(),
		Text:       text,
		NoteNature: noteNature,
	}
}

func (provider NoteProvider) Accept(e entity.EntityPick) (bool, entity.Entity) {
	if e.Nature() != noteNature {
		return false, nil
	}

	return true, &Note{}
}
