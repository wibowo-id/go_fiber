package request

type OtpRequest struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
