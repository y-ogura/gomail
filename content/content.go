package content

import (
	"os"
	"strings"
)

// Content mail content struct
type Content struct {
	From        string
	FromName    string
	To          []string
	Subject     string
	Message     string
	Attachments []Attachment
}

// Attachment mail attachment struct
type Attachment struct {
	Type     string
	FileName string
	Data     *os.File
}

// RFC5321To parse 'To' to RFC5321
func (c *Content) RFC5321To() []string {
	res := make([]string, len(c.To))
	for i, mail := range c.To {
		res[i] = formatRFC5321(mail)
	}
	return res
}

// formatRFC5321 check mail address format and format mail address to RFC5321
func formatRFC5321(mail string) string {
	mailStrings := strings.Split(mail, "@")
	if strings.LastIndex(mailStrings[0], ".") == len(mailStrings[0])-1 {
		return "\"" + mailStrings[0] + "\"@" + mailStrings[1]
	}
	return mail
}
