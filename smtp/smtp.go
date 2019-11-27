package smtp

import (
	"net/smtp"
	"os"

	mailContent "github.com/y-ogura/gomail/content"
)

// SMTP smtp struct
type SMTP struct {
	host string
	port string
	auth smtp.Auth
}

// Init initialize smtp
func Init() *SMTP {
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	user := os.Getenv("MAIL_USERNAME")
	pass := os.Getenv("MAIL_PASSWORD")

	auth := smtp.PlainAuth("", user, pass, host)
	init := &SMTP{
		host: host,
		port: port,
		auth: auth,
	}

	return init
}

// Send send email
func (s *SMTP) Send(content mailContent.Content) error {
	var msg string
	msg += "From:" + content.From + "\r\n"
	for _, to := range content.To {
		msg = msg + "To:" + to + "\r\n"
	}
	msg = msg + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg = msg + "Subject:" + content.Subject + "\r\n\r\n"
	msg = msg + content.Message

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(s.host+":"+s.port, s.auth, content.From, content.To, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
