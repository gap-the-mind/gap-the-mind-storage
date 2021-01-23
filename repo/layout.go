package repo

import (
	"fmt"
	"strings"

	"github.com/go-git/go-billy/v5"
)

var layout = map[string]string{
	"note":      "notes",
	"rendering": "rendering",
}

func id(filename string) string {
	return strings.TrimSuffix(filename, ".toml")
}

func prefix(nature string) string {
	prefix, found := layout[nature]

	if !found {
		prefix = "miscellanous"
	}

	return prefix

}

func path(fs billy.Filesystem, entity EntityRef) string {
	filename := fmt.Sprintf("%s.toml", entity.Id())
	prefix := prefix(entity.Nature())
	path := fs.Join(prefix, filename)

	return path

}
