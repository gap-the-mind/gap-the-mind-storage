//go:generate rm -rf generated
//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/gap-the-mind/gap-the-mind-storage/graph/generated"
	"github.com/gap-the-mind/gap-the-mind-storage/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user  model.User
	notes []model.Note
}

func NewResolver() generated.Config {
	r := Resolver{}

	r.user = model.User{
		ID:   "me",
		Name: "Matthieu",
	}

	r.notes = []model.Note{
		model.Note{
			ID:    "note_1",
			Title: "First note",
			Text:  "This is the first note",
		},
	}

	return generated.Config{
		Resolvers: &r,
	}
}
