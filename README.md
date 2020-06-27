# Gap the Mind Storage

## Init

    go run github.com/99designs/gqlgen init

## Generation

    go generate ./...

## Conversion to relay

    1. Generate
    1. Rename model fcrom models_gen to model
    1. Suppress all *Connection fields
    1. Generate
