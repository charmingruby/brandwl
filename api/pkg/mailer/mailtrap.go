package mailer

import (
	"github.com/go-mail/mail"
)

func NewMailTrapMailer(host string, port int, from, user, password string) *MailTrapMailer {
	return &MailTrapMailer{
		Host:     host,
		Port:     port,
		From:     from,
		User:     user,
		Password: password,
	}
}

type MailTrapMailer struct {
	Host     string
	Port     int
	From     string
	User     string
	Password string
}

func (m *MailTrapMailer) SendEmail(recipientEmail, subject, content string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", recipientEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", content)

	d := mail.NewDialer(m.Host, m.Port, m.User, m.Password)
	if err := d.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
