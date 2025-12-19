package posts

import "embed"

//go:embed *.md
var PostFS embed.FS
