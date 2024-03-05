package mail

import (
	"fmt"
	"go_fiber_wibowo/utils/config"
	"log"
	"net/smtp"
	"strings"
)

type Cfg struct {
	config *config.Config
}

func main() {
	to := []string{"recipient1@gmail.com", "emaillain@gmail.com"}
	cc := []string{"tralalala@gmail.com"}
	subject := "Test mail"
	message := "Hello"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

func sendMail(to []string, cc []string, subject, message string) error {
	_i := &config.Config{}
	body := "From: " + _i.Mail.ConfigSenderName + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", _i.Mail.ConfigAuthEmail, _i.Mail.ConfigAuthPassword, _i.Mail.ConfigSmtpHost)
	smtpAddr := fmt.Sprintf("%s:%d", _i.Mail.ConfigSmtpHost, _i.Mail.ConfigSmtpPort)

	err := smtp.SendMail(smtpAddr, auth, _i.Mail.ConfigAuthEmail, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
