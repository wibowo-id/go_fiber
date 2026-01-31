package controller

import "go_fiber_wibowo/app/module/article/service"

type Controller struct {
	Article ArticleController
}

func NewController(articleService service.ArticleService) *Controller {
	return &Controller{
		Article: NewArticleController(articleService),
	}
}
