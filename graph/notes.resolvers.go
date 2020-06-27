package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"

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

	r.notes[id] = note

	return &note, nil
}

func (r *mutationResolver) EditNote(ctx context.Context, id string, edition model.EditNoteInput) (*model.Note, error) {
	if note, found := r.notes[id]; found {
		if edition.Title != nil {
			note.Title = *edition.Title
		}

		if edition.Text != nil {
			note.Text = *edition.Text
		}

		r.notes[id] = note

		return &note, nil
	}

	return nil, fmt.Errorf("No note with ID %s", id)

}

func (r *queryResolver) CurrentUser(ctx context.Context) (*model.User, error) {
	return &r.user, nil
}

func (r *userResolver) Node(ctx context.Context, obj *model.User, id string) (model.Node, error) {
	if note, found := r.notes[id]; found {
		return note, nil
	}

	return nil, nil
}

func (r *userResolver) NotesConnection(ctx context.Context, obj *model.User, first *int, after *string, last *int, before *string) (*model.UserNotesConnection, error) {
	edges := make([]*model.UserNoteEdge, len(r.notes))

	i := 0

	for id, note := range r.notes {
		noteCopy := note

		edges[i] = &model.UserNoteEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(id)),
			Node:   &noteCopy,
		}

		i++
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
