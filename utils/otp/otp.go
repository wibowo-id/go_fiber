package otp

import (
	"dlh_oss_be/utils/helpers"
)

func GenerateOtp() string {
	Otp := helpers.RandNumberBytes(6)

	return Otp
}
