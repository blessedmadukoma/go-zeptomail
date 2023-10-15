package zeptomail

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSMTP(t *testing.T) {
	// set the .env values
	os.Setenv("SMTP_HOST", "localhost")
	os.Setenv("SMTP_PORT", "12345")
	os.Setenv("SMTP_USERNAME", "username")
	os.Setenv("SMTP_PASSWORD", "password")
	os.Setenv("SMTP_EMAIL_ADDRESS", "email@email.com")

	smtp := getSMTP()

	assert.Equal(t, smtp.Host, "localhost")
	assert.Equal(t, smtp.Port, 12345)
	assert.Equal(t, smtp.Username, "username")
	assert.Equal(t, smtp.Password, "password")
	assert.Equal(t, smtp.SenderEmail, "email@email.com")

	// unset the .env values
	os.Unsetenv("SMTP_HOST")
	os.Unsetenv("SMTP_PORT")
	os.Unsetenv("SMTP_USERNAME")
	os.Unsetenv("SMTP_PASSWORD")
	os.Unsetenv("SMTP_EMAIL_ADDRESS")
}

func TestMain(t *testing.T) {
	smtp := getSMTP()
	assert.NotNil(t, smtp)

	mailer := New(smtp)
	assert.NotNil(t, mailer)

	data := MailData{
		RecipientName:  "Blessed M.",
		RecipientEmail: "b@b.com",
		TemplateFile:   "test_template.html",
	}
	assert.NotNil(t, data)

	response, err := mailer.Send(data)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
