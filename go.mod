module github.com/gap-the-mind/gap-the-mind-storage

go 1.14

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.1.0
	github.com/google/uuid v1.1.1
	github.com/pelletier/go-toml v1.8.0
	github.com/vektah/gqlparser/v2 v2.0.1
	go.uber.org/zap v1.15.0
)

replace github.com/go-git/go-billy/v5 => ../go-billy

replace github.com/go-git/go-git/v5 => ../go-git
