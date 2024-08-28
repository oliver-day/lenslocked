package main

import (
	"os"

	"github.com/go-mail/mail/v2"
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
	msg.SetBody("text/html", html)
	msg.WriteTo(os.Stdout)
}
