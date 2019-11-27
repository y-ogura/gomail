package gomail

import (
	mailContent "github.com/y-ogura/gomail/content"
	"github.com/y-ogura/gomail/smtp"
)

// Mail mail interface
type Mail interface {
	Send(content mailContent.Content) error
}

// New create mail handler
func New(driver string) Mail {
	switch driver {
	default:
		return smtp.Init()
	}
}
