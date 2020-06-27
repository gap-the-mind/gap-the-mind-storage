package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gap-the-mind/gap-the-mind-storage/graph"
	"github.com/gap-the-mind/gap-the-mind-storage/graph/generated"
)

// StartServer starts the server
func StartServer(port string) error {
	cfg, err := graph.NewResolver()

	if err != nil {
		return err
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}
