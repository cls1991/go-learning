package packages

import "testing"

func TestMail(t *testing.T) {
	mail := newMail("user@gmail.com;user@hotmail.com", "test mail", "hello, golang")
	sendMail(mail)
}
