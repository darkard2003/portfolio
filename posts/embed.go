//go:build !dev

package posts

import (
	"embed"
	"io/fs"
)

//go:embed *.md
var postFS embed.FS

func GetPostFS() fs.FS {
	return &postFS
}
