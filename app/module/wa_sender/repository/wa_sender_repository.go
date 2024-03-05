package repository

import (
	"fmt"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/wa_sender/request"
	"go_fiber_wibowo/app/module/wa_sender/response"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/config"
	"io/ioutil"
	"net/http"
	"strings"
)

type waSenderRepository struct {
	DB         *database.Database
	Cfg        *config.Config
	OtpRequest request.OtpRequest
}

//go:generate mockgen -destination=wa_sender_repository_mock.go -package=repository . WaSenderRepository
type WaSenderRepository interface {
	SendOtp(Otp, Recipient, Message string) (user response.OtpResponse, err error)
	UpdateOtp(Otp, Phone string) (user *schema.User, err error)
}

func NewWaSenderRepository(db *database.Database, Cfg *config.Config) WaSenderRepository {
	return &waSenderRepository{
		DB:  db,
		Cfg: Cfg,
	}
}

func (_i *waSenderRepository) SendOtp(Otp, Recipient, Message string) (user response.OtpResponse, err error) {
	url := _i.Cfg.App.Wa.UrlText + "/send-message"
	method := "POST"

	payload := strings.NewReader(`{
    "api_key": "` + _i.Cfg.App.Wa.ApiKey + `",
    "sender": "` + _i.Cfg.App.Wa.Phone + `",
    "message": "*` + _i.Cfg.App.Name + `*\n\n` + Message + `Kode OTP anda adalah ` + Otp + `\n\nJangan berikan kode ini kepada siapapun !",
    "number": ` + Recipient + `
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	user.Otp = Otp
	user.Phone = Recipient

	return
}

func (_i *waSenderRepository) UpdateOtp(Otp, Phone string) (res *schema.User, err error) {
	if err := _i.DB.DB.Table("users").Where("telepon", Phone).Update("otp", Otp).Error; err != nil {
		return res, err
	}

	return res, nil
}
