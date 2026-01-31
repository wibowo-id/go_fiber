package helpers

import "strings"

func VerifyFormatPhone(phone string) string {
	if phone[0:1] == "0" {
		phone = strings.Replace(phone, "0", "62", 1)
	}

	if phone[0:2] != "62" {
		phone = "62" + phone
	}
	return phone
}
