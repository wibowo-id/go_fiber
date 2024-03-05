package controller

import "dlh_oss_be/app/module/web_page/service"

type Controller struct {
	WebPage WebPageController
}

func NewController(webPageService service.WebPageService) *Controller {
	return &Controller{
		WebPage: NewWebPageController(webPageService),
	}
}
