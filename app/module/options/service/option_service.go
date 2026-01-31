package service

import (
	"go_fiber_wibowo/app/module/options/repository"
	"go_fiber_wibowo/app/module/options/response"
)

// OptionService
type optionService struct {
	optionRepo repository.OptionRepository
}

// define interface of IOptionService
//
//go:generate mockgen -destination=option_service_mock.go -package=service . OptionService
type OptionService interface {
	GetWilayah(kode string) (res []*response.OptionsResponse, err error)
	GetRole() (res []*response.OptionsResponse, err error)
	GetPermission() (res []*response.OptionsResponse, err error)
}

// init OptionService
func NewOptionService(optionRepo repository.OptionRepository) OptionService {
	return &optionService{
		optionRepo: optionRepo,
	}
}

func (_i *optionService) GetWilayah(kode string) (res []*response.OptionsResponse, err error) {
	res, err = _i.optionRepo.GetWilayah(kode)
	if err != nil {
		return
	}

	return
}

func (_i *optionService) GetRole() (res []*response.OptionsResponse, err error) {
	res, err = _i.optionRepo.GetRole()
	if err != nil {
		return
	}

	return
}

func (_i *optionService) GetPermission() (res []*response.OptionsResponse, err error) {
	res, err = _i.optionRepo.GetPermission()
	if err != nil {
		return
	}

	return
}
