package models

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Clients = []Client{
	{Name: "Jesse Yahoo", Email: "jesseomolo@yahoo.com"},
	{Name: "Jesse Mooka", Email: "jessemooka@gmail.com"},
}

const EmailTemplate = `
Hi {{.Name}},
I hope this email finds you well.
I'm excited to share that you're eligible for a limited time offer on our website building offer. Order your website and get 50%\ off on your purchase. Take advantage of this offer before it's expiration on 31st January 2025.

Happy Holidays and a Prosperous New Year!
David Jesse,

`
