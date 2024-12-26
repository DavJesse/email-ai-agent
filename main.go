package main

import (
	"bytes"
	"log"
	"text/template"

	"email-ai-agent/ai"
	"email-ai-agent/mailhandler"
	"email-ai-agent/models"
)

func main() {

	context := `The recipient is a premium client of our company. They have recently completed a purchase, and we want to thank them and share some upcoming offers.`
	emailTemplate, err := ai.GenerateEmail(context)
	if err != nil {
		log.Fatalf("Failed to generate email: %v", err)
	}
	// Parse the email template and check for errors
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	// Iterate over the clients and send emails for each one
	for _, client := range models.Clients {
		var body bytes.Buffer

		// Execute the template with the client data and write the output to the buffer
		err = tmpl.Execute(&body, client)
		if err != nil {
			log.Printf("Failed to generate email for %s: %v", client.Name, err)
			continue
		}

		err = mailhandler.SendMail(client.Email, body.String(), "Discounted Website Building Offer")
		if err != nil {
			log.Printf("Failed to send email to %s: %v", client.Name, err)
		} else {
			log.Printf("Sent email to %s", client.Name)
		}
	}
}
