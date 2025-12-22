//go:build dev

package posts

import (
	"io/fs"
	"os"
)

func GetPostFS() fs.FS {
	return os.DirFS("posts")
}
