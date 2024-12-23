package mailhandler

import (
	"net/smtp"
)

func SendMail(to, body, subject string) error {
	var err error

	from := "jesseomolo@gmail.com"
	password := "your_gmail_password"
	smtpHost := "smtp.gmail.com"
	smtpHostPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	message := []byte("Subject: " + subject + "\r\n" + "content-type: text/plain; charset=\"utf-8\"\r\n\r\n" + body)

	err = smtp.SendMail(smtpHost+":"+smtpHostPort, auth, from, []string{to}, message)

	return err
}
