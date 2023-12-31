// Package zeptomail provides a Go client for sending emails using the ZeptoMail API.
//
// See examples and usage details on GitHub: https://github.com/blessedmadukoma/go-zeptomail
//
// To use, create a client with your API key, construct an email, and send it.
//
// Example:
//
//	smtp := zeptomail.SMTP{...}
//	client := zeptomail.New(smtp)
//	data := zeptomail.MailData{...}
//	err := client.Send(data)
package zeptomail

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type SMTP struct {
	Host        string
	Port        int
	Username    string
	Password    string
	SenderEmail string
}

type MailData struct {
	RecipientName  string
	RecipientEmail string
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
	smtp.SenderEmail = os.Getenv("SMTP_EMAIL_ADDRESS")

	return smtp
}

func main() {
	smtp := getSMTP()

	mailer := New(smtp)

	data := MailData{
		RecipientName:  "Blessed M.",
		RecipientEmail: "recipient@gmail.com",
		TemplateFile:   "welcome.html",
	}

	response, err := mailer.Send(data)
	if err != nil {
		log.Fatal("Error sending mail: ", err)
	}

	log.Println("response:", response)
}
