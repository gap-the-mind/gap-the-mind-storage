package server

import (
	"context"
	"time"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"github.com/samsarahq/thunder/reactive"
)

type post struct {
	Title     string
	Body      string
	CreatedAt time.Time
}

// server is our graphql server.
type server struct {
	posts []post
}

// registerQuery registers the root query type.
func (s *server) registerQuery(schema *schemabuilder.Schema) {
	obj := schema.Query()

	obj.FieldFunc("posts", func() []post {
		return s.posts
	})
}

// registerMutation registers the root mutation type.
func (s *server) registerMutation(schema *schemabuilder.Schema) {
	obj := schema.Mutation()
	obj.FieldFunc("echo", func(args struct{ Message string }) string {
		return args.Message
	})
}

// registerPost registers the post type.
func (s *server) registerPost(schema *schemabuilder.Schema) {
	obj := schema.Object("Post", post{})
	obj.FieldFunc("age", func(ctx context.Context, p *post) string {
		reactive.InvalidateAfter(ctx, 5*time.Second)
		return time.Since(p.CreatedAt).String()
	})
}

// schema builds the graphql schema.
func (s *server) schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	s.registerQuery(builder)
	s.registerMutation(builder)
	s.registerPost(builder)
	return builder.MustBuild()
}
