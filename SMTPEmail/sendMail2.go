package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	user := "liqinghuisi@163.com"
	password := "********"
	host := "smtp.163.com:25"
	to := "liqh@cloudhealth99.com"

	subject := "Test Golang-Email-Test"

	body := `
		<html>
		<body>
		<h3>		Test send to email		</h3>
		<p>The BCL files of following flowcells can be deleted now:</p>
		<p>CHG-S403-01A: /161221_ST-E00167_0278_AHF72VALXX; </p>
		<p>CHG-S403-01B: /161221_ST-E00167_0279_BHF2NVALXX;</p>
		<p>CHG-S403-08B: /161221_ST-E00299_0213_BHF2TVALXX</p>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}

