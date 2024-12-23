package models

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Clients = []Client{
	{Name: "Jesse Yahoo", Email: "jesseomolo@yahoo.com"},
	{Name: "Jesse Mooka", Email: "jessemooka@gmail.com"},
}
