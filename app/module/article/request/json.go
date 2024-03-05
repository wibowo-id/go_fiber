package request

import (
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/utils/paginator"
)

type ArticleRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=255"`
	Content string `json:"content" validate:"required,min=3"`
}

type ArticlesRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

func (req *ArticleRequest) ToDomain() *schema.Article {
	return &schema.Article{
		Title:   req.Title,
		Content: req.Content,
	}
}
