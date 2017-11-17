package packages

import (
	"fmt"
	"strings"
	"net/smtp"
)

const (
	HOST     = "smtp-mail.outlook.com"
	SERVER   = "smtp-mail.outlook.com:587"
	USER     = "user@hotmail.com"
	PASSWORD = "password"
)

type Mail struct {
	to      string "to"
	subject string "subject"
	msg     string "msg"
}

func NewMail(to, subject, msg string) *Mail {
	return &Mail{to, subject, msg}
}

func SendMail(mail *Mail) {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	sendTo := strings.Split(mail.to, ";")
	done := make(chan error, 1024)

	go func() {
		for _, r := range sendTo {
			c := strings.Replace("From: "+USER+"~To: "+r+"~Subject: "+mail.subject+"~~", "~", "\r\n", -1) + mail.msg
			err := smtp.SendMail(
				SERVER,
				auth,
				USER,
				[]string{r},
				[]byte(c),
			)
			if err != nil {
				done <- err
			}
		}
		close(done)
	}()
	for i := range done {
		fmt.Println(i)
	}
}
