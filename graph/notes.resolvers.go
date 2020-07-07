package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"

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

	err := r.storage.Create(NOTE_TYPE, id, note)

	return &note, err
}

func (r *mutationResolver) EditNote(ctx context.Context, id string, edition model.EditNoteInput) (*model.Note, error) {
	logger.Debugw("Edit Note",
		"id", id, "edition", edition)

	note := model.Note{}
	r.storage.Get(NOTE_TYPE, id, &note)

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

	r.storage.Update(NOTE_TYPE, id, note)

	return &note, nil
}

func (r *mutationResolver) DeleteNode(ctx context.Context, id string) (*model.Note, error) {
	note := model.Note{}
	err := r.storage.Get(NOTE_TYPE, id, &note)

	r.storage.Delete(NOTE_TYPE, id)

	return &note, err
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*model.User, error) {
	return &r.user, nil
}

func (r *userResolver) Node(ctx context.Context, obj *model.User, id string) (model.Node, error) {
	note := model.Note{}
	err := r.storage.Get(NOTE_TYPE, id, &note)

	return note, err
}

func (r *userResolver) NotesConnection(ctx context.Context, obj *model.User, first *int, after *string, last *int, before *string) (*model.UserNotesConnection, error) {
	ids, err := r.storage.List(NOTE_TYPE)

	if err != nil {
		return nil, err
	}

	edges := make([]*model.UserNoteEdge, len(ids))

	for i, id := range ids {
		note := model.Note{}

		r.storage.Get(NOTE_TYPE, id, &note)

		edges[i] = &model.UserNoteEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(id)),
			Node:   &note,
		}
	}

	return &model.UserNotesConnection{
		Edges:      edges,
		TotalCount: len(edges),
		PageInfo: &model.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
			StartCursor:     &edges[0].Cursor,
			EndCursor:       &edges[len(edges)-1].Cursor,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
