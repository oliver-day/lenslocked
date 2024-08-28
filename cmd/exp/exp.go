package main

import (
	"os"

	"github.com/go-mail/mail/v2"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "6ecc5cfb0f4a67"
	password = "2b929bf6107e7b"
)

func main() {
	from := "test@lenslocked.com"
	to := "oliver@gmail.com"
	subject := "This is the subject of an email"
	plaintext := "This is the body of an email"
	html := "<h1>Hello there buddy!</h1><p>This is an email from <b>lenslocked</b></p>"
	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}

}
