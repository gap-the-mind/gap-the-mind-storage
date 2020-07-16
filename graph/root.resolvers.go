package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gap-the-mind/gap-the-mind-storage/graph/generated"
	"github.com/gap-the-mind/gap-the-mind-storage/graph/model"
)

func (r *mutationResolver) CreateNote(ctx context.Context, title *string) (*model.Note, error) {
	note := model.Note{
		Title: *title,
	}

	return &note, r.storage.Create(&note)
}

func (r *mutationResolver) EditNote(ctx context.Context, id string, edition model.EditNoteInput) (*model.Note, error) {
	note := model.Note{
		ID: id,
	}

	r.storage.Get(&note)

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

	return &note, r.storage.Update(&note)
}

func (r *mutationResolver) DeleteNote(ctx context.Context, id string) (*model.Note, error) {
	note := model.Note{
		ID: id,
	}

	return &note, r.storage.Delete(&note)
}

func (r *mutationResolver) CreateRendering(ctx context.Context, name *string) (*model.Rendering, error) {
	rendering := model.Rendering{
		Name:  name,
		Lanes: make([]*model.Lane, 0),
	}

	return &rendering, r.storage.Create(&rendering)
}

func (r *mutationResolver) EditRendering(ctx context.Context, id string, edition model.EditRenderingInput) (*model.Rendering, error) {
	rendering := model.Rendering{
		ID: id,
	}

	r.storage.Get(&rendering)

	if edition.Lanes != nil {
		rendering.Lanes = make([]*model.Lane, len(edition.Lanes))

		for i, l := range edition.Lanes {
			rendering.Lanes[i] = &model.Lane{
				ID:     l.ID,
				Filter: l.Filter,
			}
		}
	}

	return &rendering, r.storage.Update(&rendering)
}

func (r *mutationResolver) DeleteRendering(ctx context.Context, id string) (*model.Rendering, error) {
	rendering := model.Rendering{
		ID: id,
	}

	return &rendering, r.storage.Delete(&rendering)
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
