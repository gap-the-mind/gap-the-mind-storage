//go:generate rm -rf generated
//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/gap-the-mind/gap-the-mind-storage/graph/generated"
	"github.com/gap-the-mind/gap-the-mind-storage/graph/model"
	"github.com/gap-the-mind/gap-the-mind-storage/log"
	"github.com/gap-the-mind/gap-the-mind-storage/repo"
)

var logger = log.CreateLogger()

type Resolver struct {
	user    model.User
	storage repo.Storage
}

func NewResolver() (generated.Config, error) {
	r := Resolver{}

	r.user = model.User{
		ID:    "me",
		Name:  "Matthieu",
		Email: "matthieu.dartiguenave@gmail.com",
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
