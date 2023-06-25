package smtp_server

import (
	"fmt"
	"net/smtp"
)

type smtpServer struct {
	sender         string
	senderPassword string
	server         string
	port           string
}

type smtpInterface interface {
	SendEmail(to, subject, body string) error
}

func NewServerSMTP(
	sender string,
	senderPassword string,
	server string,
	port string,
) smtpInterface {
	return &smtpServer{
		sender:         sender,
		senderPassword: senderPassword,
		server:         server,
		port:           port,
	}
}

func (s *smtpServer) SendEmail(to, subject, body string) (err error) {
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
