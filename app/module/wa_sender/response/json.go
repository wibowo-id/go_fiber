package response

type OtpResponse struct {
	Otp   string `json:"otp"`
	Phone string `json:"phone"`
}

type UpdateOtpResponse struct {
	Otp     string `json:"otp"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

type GenerateQr struct {
	Status  bool   `json:"status"`
	Qrcode  string `json:"qrcode"`
	Message string `json:"message"`
}
