package packages

import "testing"

func TestMail(t *testing.T) {
	mail := NewMail("xxx@gmail.com;xxx@hotmail.com", "test mail", "hello, golang")
	SendMail(mail)
}
