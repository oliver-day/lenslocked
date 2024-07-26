package main

import (
	"html/template"
	"os"
)

type User struct {
	Bio string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Bio: `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
