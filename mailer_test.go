package zeptomail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	smtp := SMTP{
		Host:        "smtp.example.com",
		Port:        465,
		Username:    "testuser",
		Password:    "testpassword",
		SenderEmail: "test@example.com",
	}

	mailer := New(smtp)

	assert.NotNil(t, mailer)
	assert.Equal(t, smtp.Host, mailer.dialer.Host)
	assert.Equal(t, smtp.Port, mailer.dialer.Port)
	assert.Equal(t, smtp.Username, mailer.dialer.Username)
	assert.Equal(t, smtp.Password, mailer.dialer.Password)
	assert.Equal(t, smtp.SenderEmail, mailer.sender)
}

func TestLoadTemplate(t *testing.T) {
	tmpl, err := LoadTemplate("test_template.html")

	assert.Nil(t, err)
	assert.NotNil(t, tmpl)
}
