package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// Email represents an email structure.
type Email struct {
	// SMTPInfo contains the SMTP information.
	*SMTPInfo
}

// SMTPInfo  contains the SMTP information.
type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

// NewEmail creates a new email.
func NewEmail(info *SMTPInfo) *Email {
	return &Email{
		SMTPInfo: info,
	}
}

// SendMail sends the email.
func (e Email) SendMail(to []string, subject, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("Subject", subject)
	m.SetHeader("To", to...)
	m.SetBody("text/html", body)
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
