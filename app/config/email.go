package config

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "dimasaldian9@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "dimasaldian9@gmail.com", "ppcv yfcr fvrt ouck")

	// Kirim email dan tangani error
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}
	return nil
}
