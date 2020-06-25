package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"

	"github.com/gap-the-mind/gap-the-mind-storage/graph/generated"
	"github.com/gap-the-mind/gap-the-mind-storage/graph/model"
)

func (r *queryResolver) CurrentUser(ctx context.Context) (*model.User, error) {
	return &r.user, nil
}

func (r *userResolver) NotesConnection(ctx context.Context, obj *model.User, first *int, after *string, last *int, before *string) (*model.UserNotesConnection, error) {
	edges := make([]*model.UserNoteEdge, len(r.notes))

	for i, note := range r.notes {
		edges[i] = &model.UserNoteEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(note.ID)),
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
			EndCursor:       &edges[len(edges)].Cursor,
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
