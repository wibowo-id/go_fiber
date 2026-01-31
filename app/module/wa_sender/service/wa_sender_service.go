package service

import (
	"go_fiber_wibowo/app/module/wa_sender/repository"
	"go_fiber_wibowo/app/module/wa_sender/request"
	"go_fiber_wibowo/app/module/wa_sender/response"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/helpers"
)

type waSenderService struct {
	waSenderRepo repository.WaSenderRepository
	DB           *database.Database
}

// define interface of IWaSenderService
//
//go:generate mockgen -destination=wa_sender_service_mock.go -package=service . WaSenderService
type WaSenderService interface {
	SendOtp(req request.OtpRequest) (res *response.OtpResponse, err error)
	UpdateOtp(Otp, Phone string) error
}

// init WaSenderService
func NewWaSenderService(waSenderRepo repository.WaSenderRepository, DB *database.Database) WaSenderService {
	return &waSenderService{
		waSenderRepo: waSenderRepo,
		DB:           DB,
	}
}

func (_i *waSenderService) SendOtp(req request.OtpRequest) (res *response.OtpResponse, err error) {
	Otp := helpers.RandNumberBytes(6)
	Phone := helpers.VerifyFormatPhone(req.Phone)

	user, err := _i.waSenderRepo.SendOtp(Otp, Phone, req.Message)
	if err != nil {
		return
	}

	err = _i.UpdateOtp(Otp, user.Phone)
	if err != nil {
		return res, err
	}

	return
}

func (_i *waSenderService) UpdateOtp(Otp, Phone string) error {
	res := response.UpdateOtpResponse{}
	qState := "UPDATE users "
	qState = qState + "SET otp = '" + Otp + "' "
	qState = qState + "WHERE telepon = '" + Phone + "' "

	if err := _i.DB.DB.Raw(qState).Scan(&res).Error; err != nil {
		return err
	}
	return nil
}
