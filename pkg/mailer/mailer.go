package mailer

type Mailer interface {
	SendEmail(recipientEmail, subject, content string) error
}
