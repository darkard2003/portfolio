package services

import (
	"bytes"
	"html/template"
	"io/fs"
	"log"
	"path/filepath"
	"portfolio/internals/models"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
)

type PostService struct {
	Posts     []models.PostModel
	PostCache map[string]models.PostModel
}

func NewPostService(filesystem fs.FS) (*PostService, error) {

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
	)
	var posts []models.PostModel
	cache := make(map[string]models.PostModel)

	files, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		log.Println("Error reading posts")
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		content, err := fs.ReadFile(filesystem, file.Name())
		if err != nil {
			log.Println("Eror reading post " + file.Name())
			return nil, err
		}

		var post models.PostModel
		rest, err := frontmatter.Parse(bytes.NewReader(content), &post)

		post.Slug = strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		date, err := time.Parse("02-01-2006", post.Date)
		if err != nil {
			date = time.Now()
		}

		post.DateObj = date

		if err != nil {
			log.Println("Error parsing frontmatter for file " + file.Name())
			return nil, err
		}

		var buff bytes.Buffer

		if err = md.Convert(rest, &buff); err != nil {
			log.Println("Error parsing markdown for file" + file.Name())
			return nil, err
		}

		post.Content = buff.String()
		post.ContentHTML = template.HTML(post.Content)

		posts = append(posts, post)
		if slug, ok := cache[post.Slug]; ok {
			log.Println("Duplicate slug ", slug)
			continue
		}

		cache[post.Slug] = post
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].DateObj.After(posts[j].DateObj)
	})

	return &PostService{
		Posts:     posts,
		PostCache: cache,
	}, nil

}
