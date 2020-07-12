package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"

	"github.com/gap-the-mind/gap-the-mind-storage/graph/generated"
	"github.com/gap-the-mind/gap-the-mind-storage/graph/model"
)

func (r *userResolver) Node(ctx context.Context, obj *model.User, id string) (model.Node, error) {
	note := model.Note{}
	err := r.storage.Get(Note, id, &note)

	return note, err
}

func (r *userResolver) NotesConnection(ctx context.Context, obj *model.User, first *int, after *string, last *int, before *string) (*model.UserNotesConnection, error) {
	ids := r.storage.List(Note)

	edges := make([]*model.UserNoteEdge, len(ids))

	for i, id := range ids {
		note := model.Note{}

		r.storage.Get(Note, id, &note)

		edges[i] = &model.UserNoteEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(id)),
			Node:   &note,
		}
	}

	pageInfo := model.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(edges) > 0 {
		pageInfo.StartCursor = &edges[0].Cursor
		pageInfo.EndCursor = &edges[len(edges)-1].Cursor

	}

	return &model.UserNotesConnection{
		Edges:      edges,
		TotalCount: len(edges),

		PageInfo: &pageInfo,
	}, nil
}

func (r *userResolver) RenderingsConnection(ctx context.Context, obj *model.User, first *int, after *string, last *int, before *string) (*model.UserRenderingsConnection, error) {
	ids := r.storage.List(Rendering)
	edges := make([]*model.UserRenderingEdge, len(ids))

	for i, id := range ids {
		rendering := model.Rendering{}

		r.storage.Get(Rendering, id, &rendering)

		edges[i] = &model.UserRenderingEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(id)),
			Node:   &rendering,
		}
	}

	pageInfo := model.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(edges) > 0 {
		pageInfo.StartCursor = &edges[0].Cursor
		pageInfo.EndCursor = &edges[len(edges)-1].Cursor

	}

	return &model.UserRenderingsConnection{
		Edges:      edges,
		TotalCount: len(edges),
		PageInfo:   &pageInfo,
	}, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
