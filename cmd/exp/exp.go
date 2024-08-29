package main

import (
	"fmt"

	"github.com/oliver-day/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "6ecc5cfb0f4a67"
	password = "2b929bf6107e7b"
)

func main() {
	email := models.Email{
		From:      "test@lenslocked.com",
		To:        "oliver@gmail.com",
		Subject:   "This is the subject of an email",
		Plaintext: "This is the body of an email w/ an email service",
		HTML:      "<h1>Hello there buddy!</h1><p>This is an email from <b>lenslocked</b></p>",
	}

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err := es.Send(email)
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent")
}
