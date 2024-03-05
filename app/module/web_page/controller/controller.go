package controller

import "go_fiber_wibowo/app/module/web_page/service"

type Controller struct {
	WebPage WebPageController
}

func NewController(webPageService service.WebPageService) *Controller {
	return &Controller{
		WebPage: NewWebPageController(webPageService),
	}
}
