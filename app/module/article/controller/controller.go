package controller

import "dlh_oss_be/app/module/article/service"

type Controller struct {
	Article ArticleController
}

func NewController(articleService service.ArticleService) *Controller {
	return &Controller{
		Article: NewArticleController(articleService),
	}
}
