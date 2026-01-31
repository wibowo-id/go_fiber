package controller

import (
	"go_fiber_wibowo/app/module/wa_sender/service"
)

type Controller struct {
	WaSender WaSenderController
}

func NewController(waSenderService service.WaSenderService) *Controller {
	return &Controller{
		WaSender: NewWaSenderController(waSenderService),
	}
}
