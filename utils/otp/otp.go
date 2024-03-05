package otp

import (
	"go_fiber_wibowo/utils/helpers"
)

func GenerateOtp() string {
	Otp := helpers.RandNumberBytes(6)

	return Otp
}
