package articles

import "github.com/rezaAmiri123/library/conf"

type ArticleSerializer struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Body        string `json:"body"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (as *ArticleSerializer) Response(a *Article) *ArticleSerializer {
	as.ID = a.ID
	as.Title = a.Title
	as.Slug = a.Slug
	as.Description = a.Description
	as.Body = a.Body
	as.CreatedAt = a.CreatedAt.UTC().Format(conf.UTCFormat)
	as.UpdatedAt = a.UpdatedAt.UTC().Format(conf.UTCFormat)
	return as
}
