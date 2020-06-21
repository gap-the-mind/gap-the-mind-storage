package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-git/go-git/v5"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"gitlab.com/ekai/proudhon/gap-the-mind-storage/log"
)

var logger = log.CreateLogger()

func createSchema() *graphql.Schema {
	// Read and parse the schema:
	bstr, err := ioutil.ReadFile("server/proudhon.graphql")

	if err != nil {
		logger.Panicw("Failed to read schema", "error", err)
	}

	schemaString := string(bstr)

	schema, err := graphql.ParseSchema(schemaString, &RootResolver{})

	if err != nil {
		logger.Panicw("Failed to parse schema", "error", err)
	}

	return schema
}

// StartServer starts the server
func StartServer(repo *git.Repository) {
	defer logger.Sync()

	schema := createSchema()
	port := 8080

	logger.Infow("Start server",
		"schema", &graphql.Schema{},
		"port",
		port)

	http.Handle("/graphql", &relay.Handler{Schema: schema})

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	if err != nil {
		logger.Fatalw("Failed to start server", "error", err)
	}

}
