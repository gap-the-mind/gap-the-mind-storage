//go:generate rm -rf generated
//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/gap-the-mind/gap-the-mind-storage/graph/generated"
	"github.com/gap-the-mind/gap-the-mind-storage/graph/model"
	"github.com/gap-the-mind/gap-the-mind-storage/repo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user    model.User
	storage repo.Storage
}

func NewResolver() (generated.Config, error) {
	r := Resolver{}

	r.user = model.User{
		ID:   "me",
		Name: "Matthieu",
	}

	storage, err := repo.Open("../storage")

	if err != nil {
		return generated.Config{}, err
	}

	r.storage = storage

	return generated.Config{
		Resolvers: &r,
	}, nil
}
