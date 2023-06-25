package smtp_gmail

import (
	"fmt"
	"net/smtp"
)

type gmailSmtp struct {
	sender         string
	senderPassword string
	server         string
	port           string
}

type gmailSmtpInterface interface {
	SendEmail(to, subject, body string) error
}

func NewGmailSMTP(
	sender string,
	senderPassword string,
	server string,
	port string,
) gmailSmtpInterface {
	return &gmailSmtp{
		sender:         sender,
		senderPassword: senderPassword,
		server:         server,
		port:           port,
	}
}

func (s *gmailSmtp) SendEmail(to, subject, body string) (err error) {
	var auth smtp.Auth
	fmt.Println(s.sender, s.senderPassword, s.server)
	auth = smtp.PlainAuth("", s.sender, s.senderPassword, s.server)
	message := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", s.sender, to, subject, body)

	err = smtp.SendMail(s.server+":"+s.port, auth, s.sender, []string{to}, []byte(message))
	return
}
