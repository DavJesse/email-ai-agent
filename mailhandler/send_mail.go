package mailhandler

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(to, body, subject string) error {
	var err error
	// Load the .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set up the SMTP client
	from := "jesseomolo@gmail.com"
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpHostPort := "587"

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Set up the email message
	message := []byte("Subject: " + subject + "\r\n" + "content-type: text/plain; charset=\"utf-8\"\r\n\r\n" + body)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpHostPort, auth, from, []string{to}, message)

	return err
}
