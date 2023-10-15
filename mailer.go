package zeptomail

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
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

// LoadTemplate loads a template from the user's local 'templates' directory.
func LoadTemplate(fileName string) (*template.Template, error) {
	// Get the path to the user's 'templates' directory
	templatesDir := "templates"

	// Build the full path to the template file
	templatePath := filepath.Join(templatesDir, fileName)

	// Open and parse the template file
	tmpl, err := template.New("email").ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

// Send() takes a data containing the recipient email address, file name containing the templates, and any dynamic data for the templates
func (m Mailer) Send(data MailData) (message string, err error) {
	tmpl, err := LoadTemplate(data.TemplateFile)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %v", err)
	}

	tmpl.ParseFiles(data.TemplateFile)

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return "", fmt.Errorf("error executing template with subject: %v", err)
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return "", fmt.Errorf("error executing plainbody: %v", err)
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return "", fmt.Errorf("error executing HTML body: %v", err)
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", data.RecipientEmail)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(msg)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("mail sent to %s, %s!\n", data.RecipientName, data.RecipientEmail), nil
}
