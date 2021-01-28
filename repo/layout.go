package repo

import (
	"fmt"
	"github.com/gap-the-mind/gap-the-mind-storage/entity"
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
		prefix = "miscellaneous"
	}

	return prefix

}

func entityPath(fs billy.Filesystem, entity entity.Entity) string {
	filename := fmt.Sprintf("%s.json", entity.Id())
	prefix := prefix(entity.Nature())
	path := fs.Join(prefix, filename)

	return path

}
