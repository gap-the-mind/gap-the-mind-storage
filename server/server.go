package server

import (
	"net/http"

	"github.com/go-git/go-git/v5"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/graphiql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"gitlab.com/ekai/proudhon/gap-the-mind-storage/log"
)

var logger = log.CreateLogger()

// StartServer starts the server
func StartServer(repo *git.Repository) {
	defer logger.Sync()

	server := &server{}

	schema := server.schema()
	introspection.AddIntrospectionToSchema(schema)

	// Expose schema and graphiql.
	http.Handle("/graphql", graphql.Handler(schema))
	http.Handle("/graphiql/", http.StripPrefix("/graphiql/", graphiql.Handler()))

	http.ListenAndServe(":3030", nil)

}
