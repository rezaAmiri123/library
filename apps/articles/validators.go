package articles

import "github.com/gosimple/slug"

type ArticleValidator struct {
	Title       string `json:"title" binding:"required,min=4"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func (av *ArticleValidator) SetData(a *Article) error {
	a.Title = av.Title
	a.Description = av.Description
	a.Body = av.Body
	a.Slug = slug.Make(av.Title)
	return nil
}
