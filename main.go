package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type SMTP struct {
	host     string
	port     int
	username string
	password string
	sender   string
}

// getSMTP retreives the SMTP details from the env
func getSMTP() SMTP {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	var smtp SMTP

	smtp.host = os.Getenv("SMTP_HOST")
	smtp.port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtp.username = os.Getenv("SMTP_USERNAME")
	smtp.password = os.Getenv("SMTP_PASSWORD")
	smtp.sender = os.Getenv("SMTP_EMAIL_ADDRESS")

	return smtp
}

func main() {
	smtp := getSMTP()

	mailer := New(smtp.host, smtp.port, smtp.username, smtp.password, smtp.sender)

	fmt.Println(smtp.host, smtp.port, smtp.username, smtp.password, smtp.sender)

	data := map[string]string{
		"Name": "John Doe",
	}

	recepient := "receipent@mail.com"

	err := mailer.Send(recepient, "welcome.html", data)
	if err != nil {
		log.Fatal("Error sending mail: ", err)
	}
}
