// package zeptomail
package main

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"github.com/go-mail/mail/v2"
)

type Mailer struct {
	dialer *mail.Dialer
	sender string
}

// New initializes a new mail.Dialer instace
func New(smtp SMTP) Mailer {

	host, port, username, password, sender := smtp.Host, smtp.Port, smtp.Username, smtp.Password, smtp.SenderEmail

	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 5 * time.Second

	return Mailer{
		dialer: dialer,
		sender: sender,
	}
}

// Send() takes a data containing the recipient email address, file name containing the templates, and any dynamic data for the templates
func (m Mailer) Send(data MailData) error {
	// tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+data.TemplateFile)
	tmpl, err := template.New("email").ParseGlob("templates/*")
	if err != nil {
		return err
	}

	tmpl.ParseFiles(data.TemplateFile)

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", data.RecipientEmail)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	log.Println("=> Mail sent!")
	return nil
}
