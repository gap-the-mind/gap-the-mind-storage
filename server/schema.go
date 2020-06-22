package server

import (
	"fmt"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

type note struct {
	Title string
	Text  string
}

// server is our graphql server.
type server struct {
}

// registerQuery registers the root query type.
func (s *server) registerQuery(schema *schemabuilder.Schema) {
	query := schema.Query()

	query.FieldFunc("notes", func() []note {
		return []note{
			{Title: "first post!", Text: "I was here first!"},
			{Title: "graphql", Text: "did you hear about Thunder?"},
		}
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
// func (s *server) registerPost(schema *schemabuilder.Schema) {
// 	obj := schema.Object("Post", post{})
// 	obj.FieldFunc("age", func(ctx context.Context, p *post) string {
// 		reactive.InvalidateAfter(ctx, 5*time.Second)
// 		return time.Since(p.CreatedAt).String()
// 	})
// }

// schema builds the graphql schema.
func (s *server) schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	s.registerQuery(builder)
	s.registerMutation(builder)
	// s.registerPost(builder)

	valueJSON, err := introspection.ComputeSchemaJSON(*builder)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(valueJSON))

	return builder.MustBuild()
}
