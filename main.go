package main

import (
	"bytes"
	"log"
	"text/template"

	"email-ai-agent/mailhandler"
	"email-ai-agent/models"
)

func main() {
	tmpl, err := template.New("email").Parse(models.EmailTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	for _, client := range models.Clients {
		var body bytes.Buffer

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
