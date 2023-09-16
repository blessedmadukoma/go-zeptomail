package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type SMTP struct {
	Host     string
	Port     int
	Username string
	Password string
	Sender   string
}

type MailData struct {
	RecepientName  string
	RecepientEmail string
	TemplateFile   string
}

// getSMTP retreives the SMTP details from the env
func getSMTP() SMTP {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	var smtp SMTP

	smtp.Host = os.Getenv("SMTP_HOST")
	smtp.Port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtp.Username = os.Getenv("SMTP_USERNAME")
	smtp.Password = os.Getenv("SMTP_PASSWORD")
	smtp.Sender = os.Getenv("SMTP_EMAIL_ADDRESS")

	return smtp
}

func main() {
	smtp := getSMTP()

	mailer := New(smtp)

	data := MailData{
		RecepientName:  "Blessed M.",
		RecepientEmail: "recepient@gmail.com",
		TemplateFile:   "welcome.html",
	}

	err := mailer.Send(data)
	if err != nil {
		log.Fatal("Error sending mail: ", err)
	}
}
