package controller

import "go_fiber_wibowo/app/module/version/service"

type Controller struct {
	Version VersionController
}

func NewController(versionService service.VersionService) *Controller {
	return &Controller{
		Version: NewVersionController(versionService),
	}
}
