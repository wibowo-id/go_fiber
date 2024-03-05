package controller

import (
	"dlh_oss_be/app/module/wa_sender/service"
)

type Controller struct {
	WaSender WaSenderController
}

func NewController(waSenderService service.WaSenderService) *Controller {
	return &Controller{
		WaSender: NewWaSenderController(waSenderService),
	}
}
