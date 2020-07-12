package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gap-the-mind/gap-the-mind-storage/graph/generated"
	"github.com/gap-the-mind/gap-the-mind-storage/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateNote(ctx context.Context, title *string) (*model.Note, error) {
	id := uuid.New().String()

	note := model.Note{
		ID:    id,
		Title: *title,
	}

	err := r.storage.Create(Note, id, note)

	return &note, err
}

func (r *mutationResolver) EditNote(ctx context.Context, id string, edition model.EditNoteInput) (*model.Note, error) {
	logger.Debugw("Edit Note",
		"id", id, "edition", edition)

	note := model.Note{}
	r.storage.Get(Note, id, &note)

	if edition.Title != nil {
		note.Title = *edition.Title
	}

	if edition.Text != nil {
		note.Text = *edition.Text
	}

	if edition.Tags != nil {
		note.Tags = make([]*model.Tag, len(edition.Tags))

		for i, t := range edition.Tags {
			note.Tags[i] = &model.Tag{ID: t.ID}
		}

	}

	r.storage.Update(Note, id, note)

	return &note, nil
}

func (r *mutationResolver) DeleteNote(ctx context.Context, id string) (*model.Note, error) {
	note := model.Note{}
	err := r.storage.Get(Note, id, &note)

	r.storage.Delete(Note, id)

	return &note, err
}

func (r *mutationResolver) CreateRendering(ctx context.Context, name *string) (*model.Rendering, error) {
	id := uuid.New().String()

	rendering := model.Rendering{
		ID:    id,
		Name:  name,
		Lanes: make([]*model.Lane, 0),
	}

	err := r.storage.Create(Rendering, id, rendering)

	return &rendering, err
}

func (r *mutationResolver) EditRendering(ctx context.Context, id string, edition model.EditRenderingInput) (*model.Rendering, error) {
	rendering := model.Rendering{}
	r.storage.Get(Rendering, id, &rendering)

	if edition.Lanes != nil {
		rendering.Lanes = make([]*model.Lane, len(edition.Lanes))

		for i, l := range edition.Lanes {
			rendering.Lanes[i] = &model.Lane{
				ID:     l.ID,
				Filter: l.Filter,
			}
		}
	}

	r.storage.Update(Rendering, id, rendering)

	return &rendering, nil
}

func (r *mutationResolver) DeleteRendering(ctx context.Context, id string) (*model.Rendering, error) {
	rendering := model.Rendering{}
	err := r.storage.Get(Rendering, id, &rendering)

	r.storage.Delete(Rendering, id)

	return &rendering, err
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*model.User, error) {
	return &r.user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
