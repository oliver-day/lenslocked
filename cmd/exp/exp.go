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
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err := es.ForgotPassword("oliver@sublime.io", "https://lenslocked.com/reset-pw?token=abc123")
	if err != nil {
		panic(err)
	}

	fmt.Println("Email sent")
}
