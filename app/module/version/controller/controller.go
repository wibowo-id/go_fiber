package controller

import "dlh_oss_be/app/module/version/service"

type Controller struct {
	Version VersionController
}

func NewController(versionService service.VersionService) *Controller {
	return &Controller{
		Version: NewVersionController(versionService),
	}
}
